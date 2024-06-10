#!/bin/bash

randompwd()
{
    LEN=$1
    declare -a CHARS
    #CHARS=(a b c d e f g h i j k l m n o p q r s t u v w x y z A B C D E F G H I J K L M N O P Q R S T U V W X Y Z 0 1 2 3 4 5 6 7 8 9 \; \: \. \~ \! \@ \# \$ \% \^ \& \* - + = \?)
    CHARS=(a b c d e f g h i j k l m n o p q r s t u v w x y z A B C D E F G H I J K L M N O P Q R S T U V W X Y Z 0 1 2 3 4 5 6 7 8 9 \_ \. \/ \~ + -)
    CNUM="${#CHARS[@]}"
    RANDSTR=""
    POSCNT=0;

    while [ "$POSCNT" -lt "$LEN" ]
    do
	POSCNT=$((POSCNT+1))
	IDX=$(($RANDOM % $CNUM))
	RANDSTR="${RANDSTR}${CHARS[$IDX]}"
    done

    printf "%s"  "${RANDSTR}"
    return 0
}

