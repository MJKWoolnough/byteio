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

// File automatically generated with ./gen.sh.

import (
	"io"
	"math"
)

// ${s}${e}Endian${rw}${er} wraps a io.${rw}${er} to provide methods
// to make it easier to $rw fundamental types.
type ${s}${e}Endian${rw}${er} struct {
	io.${rw}${er}
	buffer [9]byte
HEREDOC
				if [ ! -z "$s" ]; then
					echo "	Err    error";
					echo "	Count  int64";
				fi;

				echo "}";

				if [ ! -z "$s" ]; then
					cat <<HEREDOC

// ${rw} implements the io.${rw}${er} interface.
func (e *${s}${e}Endian${rw}${er}) ${rw}(p []byte) (int, error) {
	if e.Err != nil {
		return 0, e.Err
	}

	var n int
HEREDOC
					if [ "$rw" = "Read" ]; then
						echo "	n, e.Err = io.ReadFull(e.Reader, p)";
					else
						echo "	n, e.Err = e.Writer.Write(p)";
					fi;
					cat <<HEREDOC
	e.Count += int64(n)

	return n, e.Err
}

HEREDOC
				else
					echo;
				fi;

				echo "// ${rw}Bool ${rw}s a boolean.";
				echo -n "func (e *${s}${e}Endian${rw}${er}) ${rw}Bool(";
				if [ "${rw}" = "Write" ]; then
					echo -n "b bool";
					ret="";
					if [ -z "$s" ]; then
						echo -n ") (int, error"
						ret="return ";
					fi;
					echo ") {"
					cat <<HEREDOC
	if b {
		${ret}e.WriteUint8(1)
	}

	${ret}e.WriteUint8(0)
}
HEREDOC
				else
					echo -n ") "
					if [ -z "$s" ]; then
						echo "(bool, int, error) {"
						echo "	b, n, err := e.ReadUint8()";
						echo;
						echo "	return b != 0, n, err";
						echo "}";
					else
						echo "bool {"
						echo "	return e.ReadUint8() != 0";
						echo "}";
					fi;
				fi;

				for t in "Int" "Uint" "Float"; do
					types="8 16 24 32 40 48 56 64"
					if [ "$t" = "Float" ]; then
						types="32 64";
					fi;
					for i in $types; do
						ti=$i;
						if [ $i -eq 24 ]; then
							ti=32;
						elif [ $i -gt 32 -a $i -lt 64 ]; then
							ti=64;
						fi;
						tu="$(echo "$t" | tr A-Z a-z)$ti";
						echo;
						echo "// ${rw}${t}${i} ${rw}s a $i bit $(echo "$t" | tr A-Z a-z) as a $tu using the underlying io.${rw}${er}.";
						if [ ! -z "$s" ]; then
							echo "// Any errors and the running byte read count are stored instead of being returned.";
						fi;
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
							echo;
						fi;
						if [ "$rw" = "Read" ]; then
							echo "	n, err := io.ReadFull(e.Reader, e.buffer[:$(( $i / 8 ))])";
							if [ ! -z "$s" ]; then
								echo;
								echo "	e.Count += int64(n)";
								echo;
							fi;
							echo "	if err != nil {";
							if [ -z "$s" ]; then
								echo "		return 0, n, err";
							else
								echo "		e.Err = err";
								echo;
								echo "		return 0";
							fi;
					
							echo "	}";
							echo ;
							echo -n "	return ";

							if [ "$t" = "Int" ]; then
								echo -n "int$ti(";
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
									echo -n "uint$ti(";
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
							if [ $i -eq 8 ]; then
								if [ "$t" = "Int" ]; then
									echo "	e.buffer[0] = byte(d)";
								else
									echo "	e.buffer[0] = d";
								fi;
							else
								if [ "$t" = "Float" ]; then
									var="c";
									echo "	c := math.Float${i}bits(d)";
								elif [ "$t" = "Int" ]; then
									var="c";
									echo "	c := uint${ti}(d)";
								fi;
								echo "	e.buffer = [9]byte{";

								shift=0;

								if [ $order -eq -1 ]; then
									shift=$(( ((i / 8) - 1) * 8));
								fi;

								for n in $(seq $(( i / 8 ))); do
									echo -n "		byte(";
									echo -n "$var";
									if [ $shift -ne 0 ]; then
										echo -n " >> $shift";
									fi;
									echo "),";
									let "shift += 8 * $order";
								done;

								echo "	}";
							fi;

							if [ -z "$s" ]; then
								echo;
								echo -n "	return ";
							else
								echo;
								echo "	var n int";
								echo;
								echo -n "	n, e.Err = ";
							fi;
							
							echo  "e.Writer.Write(e.buffer[:$(( $i / 8 ))])";
							
							if [ ! -z "$s" ]; then
								echo "	e.Count += int64(n)";
							fi;
						fi;

						echo "}";
					done;
				done;
				type="[]byte";
				for t in Bytes String; do
					tt="$t";
					if [ "$tt" = "Bytes" ]; then
						tt="";
					fi;
					if [ "$t" = "String" ] || [ "$rw" = "Read" ]; then
						echo;
						echo "// ${rw}${t} ${rw}s a ${type}.";
						echo -n "func (e *${s}${e}Endian${rw}${er}) ${rw}${t}(";
						if [ "$rw" = "Write" ]; then
							echo "d ${type}) (int, error) {";
							if [ -z "$s" ]; then
								if [ "$t" = "String" ]; then
									echo "	return io.WriteString(e.Writer, d)";
								else
									echo "	return e.Write(d)";
								fi;
							else
								echo "	if e.Err != nil {";
								echo "		return 0, e.Err";
								echo "	}";
								echo;
								echo "	var n int";
								if [ "$t" = "String" ]; then
									echo "	n, e.Err = io.WriteString(e.Writer, d)";
								else
									echo "	n, e.Err = e.Write(d)";
								fi;
								echo "	e.Count += int64(n)";
								echo;
								echo "	return n, e.Err";
							fi;
						else
							echo -n "size int) ";
							if [ -z "$s" ]; then
								echo "(${type}, int, error) {";
								echo "	buf := make([]byte, size)";
								echo "	n, err := io.ReadFull(e, buf)";
								echo;
								if [ "$t" = "String" ]; then
									echo "	return string(buf[:n]), n, err";
								else
									echo "	return buf[:n], n, err";
								fi;
							else
								echo "${type} {";
								echo "	if e.Err != nil {";
								if [ "$t" = "String" ]; then
									echo "		return \"\"";
								else
									echo "		return nil";
								fi;
								echo "	}";
								echo;
								echo "	buf := make([]byte, size)";
								echo;
								echo "	var n int";
								echo "	n, e.Err = io.ReadFull(e.Reader, buf)";
								echo "	e.Count += int64(n)";
								echo;
								if [ "$t" = "String" ]; then
									echo "	return string(buf[:n])";
								else
									echo "	return buf[:n]";
								fi;
							fi;
						fi;
						echo "}";
					fi;
					for size in "X" 8 16 24 32 40 48 56 64; do
						tSize="$size";
						if [ "$size" = "X" ]; then
							tSize="64";
						elif [ $size -eq 24 ]; then
							tSize=32;
						elif [ $size -gt 32 -a $size -lt 64 ]; then
							tSize=64;
						fi;
						echo;
						echo "// ${rw}${t}${size} ${rw}s the length of the ${t}, using ReadUint${size} and then ${rw}s the bytes.";
						echo -n "func (e *${s}${e}Endian${rw}${er}) ${rw}${t}${size}(";
						if [ "$rw" = "Write" ]; then
							if [ -z "$s" ]; then
								echo "p ${type}) (int, error) {";
								echo "	n, err := e.WriteUint${size}(uint${tSize}(len(p)))";
								echo "	if err != nil {";
								echo "		return n, err";
								echo "	}";
								echo;
								echo "	m, err := e.Write${tt}(p)";
								echo;
								echo "	return n + m, err";
							else
								echo "p ${type}) {";
								echo "	e.WriteUint${size}(uint${tSize}(len(p)))";
								echo "	e.Write${tt}(p)"
							fi;
						else
							if [ -z "$s" ]; then
								echo ") (${type}, int, error) {";
								echo "	size, n, err := e.ReadUint${size}()";
								echo "	if err != nil {"
								if [ "$t" = "String" ]; then
									echo "		return \"\", n, err";
								else
									echo "		return nil, n, err";
								fi;
								echo "	}";
								echo;
								echo "	p, m, err := e.Read${t}(int(size))"
								echo;
								echo "	return p, n + m, err";
							else
								echo ") ${type} {";
								echo "	return e.Read${t}(int(e.ReadUint${size}()))";
							fi;
						fi;
						echo "}";
					done;
					type="string";
				done;
				echo;
				echo -n "// ${rw}String0 ${rw}s the bytes of the string ";
				if [ "$rw" = "Write" ]; then
					echo "ending with a 0 byte.";
				else
					echo "until a 0 byte is read.";
				fi;
				echo -n "func (e *${s}${e}Endian${rw}${er}) ${rw}String0(";
				if [ "$rw" = "Write" ]; then
					if [ -z "$s" ]; then
						echo "p string) (int, error) {";
						echo "	n, err := e.WriteString(p)";
						echo "	if err == nil {";
						echo "		var m int";
						echo "		m, err = e.WriteUint8(0)";
						echo "		n += m";
						echo "	}";
						echo;
						echo "	return n, err";
					else
						echo "p string) {";
						echo "	e.WriteString(p)";
						echo "	e.WriteUint8(0)";
					fi;
				else
					if [ -z "$s" ]; then
						echo ") (string, int, error) {";
						echo "	var (";
						echo "		n   int";
						echo "		err error";
						echo "		d   []byte";
						echo "	)";
						echo;
						echo "	for {";
						echo "		var (";
						echo "			m int";
						echo "			p byte";
						echo "		)";
						echo;
						echo "		p, m, err = e.ReadUint8()";
						echo "		n += m";
						echo;
						echo "		if err != nil || p == 0 {";
						echo "			break";
						echo "		}";
						echo;
						echo "		d = append(d, p)";
						echo "	}";
						echo;
						echo "	return string(d), n, err";
					else
						echo ") string {";
						echo "	var d []byte";
						echo;
						echo "	for {";
						echo "		p := e.ReadUint8()";
						echo "		if p == 0 {";
						echo "			break";
						echo "		}";
						echo;
						echo "		d = append(d, p)";
						echo "	}";
						echo;
						echo "	return string(d)";
					fi;
				fi;
				echo "}";
			) > "$(echo "${s}${e}Endian${rw}${er}" | tr A-Z a-z).go";
		done;
	done;
done;
