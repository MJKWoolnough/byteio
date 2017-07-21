#!/bin/bash

for s in "" Sticky; do
	for e in Big Little; do
		declare startP;
		declare order;
		if [ "$e" = "Big" ]; then
			startP=1;
			order=-1;
		else
			startP=0;
			order=1;
		fi;

		for rw in Read Write; do
			er="er";
			if [ "$rw" = "Write" ]; then
				er="r";
			fi;
			(
				cat <<HEREDOC
package byteio

// File automatically generated with ./gen.sh

import (
	"io"
	"math"
)

type ${s}${e}Endian${rw}${er} struct {
	io.${rw}${er}
	buffer [8]byte
HEREDOC
				if [ ! -z "$s" ]; then
					echo "	Err    error";
					echo "	Count  int64";
				fi;

				echo "}";

				if [ ! -z "$s" ]; then
					cat <<HEREDOC

func (e *${s}${e}Endian${rw}${er}) ${rw}(p []byte) (int, error) {
	if e.Err != nil {
		return 0, e.Err
	}
	var n int
	n, e.Err = e.${rw}${er}.$rw(p)
	e.Count += int64(n)
	return n, e.Err
}
HEREDOC
				fi;

				for t in "Int" "Uint" "Float"; do
					types="8 16 32 64"
					if [ "$t" = "Float" ]; then
						types="32 64";
					fi;
					for i in $types; do
						tu="$(echo "$t" | tr A-Z a-z)$i";
						echo;
						echo -n "func (e *${s}${e}Endian${rw}${er}) ${rw}${t}${i}(";
						if [ "$rw" = "Write" ]; then
							echo -n "d $tu) ";
							if [ -z "$s" ]; then
								echo -n "(";
							fi;
						else
							echo -n ") ";
							if [ -z "$s" ]; then
								echo -n "($tu, ";
							else
								echo -n "$tu ";
							fi;
						fi;
						if [ -z "$s" ]; then
							echo "int, error) {";
						else
							echo "{";
						fi;

						if [ ! -z "$s" ]; then
							echo "	if e.Err != nil {";
							echo -n "		return";
							if [ "$rw" = "Read" ]; then
								echo -n " 0";
							fi;
							echo;
							echo "	}";
						fi;
						if [ "$rw" = "Read" ]; then
							echo "	n, err := io.ReadFull(e.Reader, e.buffer[:$(( $i / 8 ))])";
							if [ ! -z "$s" ]; then
								echo "	e.Count += int64(n)";
							fi;
							echo "	if err != nil {";
							if [ -z "$s" ]; then
								echo "		return 0, n, err";
							else
								echo "		e.Err = err";
								echo "		return 0";
							fi;
					
							echo "	}";
							echo -n "	return ";

							if [ "$t" = "Int" ]; then
								echo -n "int$i(";
							elif [ "$t" = "Float" ]; then
								echo -n "math.Float${i}frombits(";
							fi;
							
							p=$(( $startP * $i / 8 - startP));
							shift=0;

							for n in $(seq 1 $(( $i / 8 ))); do
								if [ $n -ne 1 ]; then
									echo -n " | ";
								fi;
								if [ "$i" -ne 8 ]; then
									echo -n "uint$i(";
								fi;
								echo -n "e.buffer[$p]";
								if [ "$i" -ne 8 ]; then
									echo -n ")";
								fi;
								if [ $shift -gt 0 ]; then
									echo -n "<<$shift";
								fi;
								let "shift += 8";
								let "p += order";
							done;

							if [ "$t" != "Uint" ]; then
								echo -n ")";
							fi;

							if [ -z "$s" ]; then
								echo ", $(( $i / 8 )), nil";
							else
								echo;
							fi;
						else
							var="d";
							if [ "$t" = "Int" -a $i -ne 8 ]; then
								var="c";
								echo "	c := uint$i(d)";
							elif [ "$t" = "Float" ]; then
								var="c";
								echo "	c := math.Float${i}bits(d)";
							fi;
							echo "	e.buffer = [8]byte{";

							shift=$(( $startP * ($i - 8) ))


							for n in $(seq 1 $(( $i / 8 ))); do
								echo -n "		";
								if [ $tu != "uint8" ]; then
									echo -n "byte(";
								fi;
								echo -n "$var";
								if [ $shift -ne 0 ]; then
									echo -n " >> $shift";
								fi;
								if [ $tu != "uint8" ]; then
									echo -n ")";
								fi;
								echo ",";
								let "shift += 8 * $order";
							done;

							echo "	}";

							if [ -z "$s" ]; then
								echo -n "	return ";
							else
								echo "	var n int";
								echo -n "	n, e.Err = ";
							fi;
							
							echo "e.Writer.Write(e.buffer[:$(( $i / 8 ))])";
							
							if [ ! -z "$s" ]; then
								echo "	e.Count += int64(n)";
							fi;
						fi;

						echo "}";
					done;
				done;
			) > "$(echo "${s}${e}Endian${rw}${er}" | tr A-Z a-z).go";
		done;
	done;
done;
