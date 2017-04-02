<?php

declare(strict_types=1);

class Task3 {
    // the length of S is within the range [0..50,000];i
    // string S consists only of the following characters: "A", "B" and/or "C".
    public function solution($input):string{

        $rules = $this->getPossibleRules($input);

        if (empty($rules)){
            return $input;
        }

        $input = $this->changeInput($input,$rules);
        return $this->solution($input);
    }

    public function solution2($input) :string {

        while (!empty($this->getPossibleRules($input))){
            $rules = $this->getPossibleRules($input);
            $input = $this->changeInput($input,$rules);
        }

        return $input;
    }


    // Assuming the input is string and only has A,B,C char
    public function getPossibleRules($input):array{
        $result = array();
        $lengthInput = strlen($input);

        // getting the rules by looping through the input variable and check it one by one.
        for ($i=0;$i<$lengthInput;$i++) {
            // avoid index out of bound
            if ($i != $lengthInput-1) {
                $subs = $input[$i] . $input[$i+1];
                switch ($subs) {
                case $subs == "AB":
                    // append with the rules number 1
                    array_push($result,1);
                    break;
                case $subs == "BA":
                    // append with the rules number 2
                    array_push($result,2);
                    break;
                case $subs == "CB":
                    // append with the rules number 3
                    array_push($result,3);
                    break;
                case $subs == "BC":
                    // append with the rules number 4
                    array_push($result,4);
                    break;
                case $subs == "AA":
                    // append with the rules number 5
                    array_push($result,5);
                    break;
                case $subs == "CC":
                    // append with the rules number 6
                    array_push($result,6);
                    break;
                default :
                    continue;

                }
            }
        }
        return $result;
    }

    public function changeInput($input, $rules):string{
        $result = str_split($input);

        // pick random rule 
        $randomKey = array_rand($rules,1);
        $randomRule = $rules[$randomKey];

        // change the input to the specific rule that picked by randomRule.
        switch ($randomRule){
        case $randomRule == 1 :
            $lengthInput = strlen($input);

            // getting the rules by looping through the input variable and check it one by one.
            for ($i=0;$i<$lengthInput;$i++) {
                // avoid index out of bound
                if ($i != $lengthInput-1) {
                    $subs = $input[$i] . $input[$i+1];
                    if ($subs == "AB"){
                        $result[$i] = "A";
                        $result[$i+1] = "A";
                        break;
                    }
                }

            }

            break;

        case $randomRule == 2 :
            $lengthInput = strlen($input);

            // getting the rules by looping through the input variable and check it one by one.
            for ($i=0;$i<$lengthInput;$i++) {
                // avoid index out of bound
                if ($i != $lengthInput-1) {
                    $subs = $input[$i] . $input[$i+1];
                    if ($subs == "BA"){
                        $result[$i] = "A";
                        $result[$i+1] = "A";
                        break;
                    }
                }

            }

            break;

        case $randomRule == 3 :
            $lengthInput = strlen($input);

            // getting the rules by looping through the input variable and check it one by one.
            for ($i=0;$i<$lengthInput;$i++) {
                // avoid index out of bound
                if ($i != $lengthInput-1) {
                    $subs = $input[$i] . $input[$i+1];
                    if ($subs == "CB"){
                        $result[$i] = "C";
                        $result[$i+1] = "C";
                        break;
                    }
                }

            }

            break;

        case $randomRule == 4 :
            $lengthInput = strlen($input);

            // getting the rules by looping through the input variable and check it one by one.
            for ($i=0;$i<$lengthInput;$i++) {
                // avoid index out of bound
                if ($i != $lengthInput-1) {
                    $subs = $input[$i] . $input[$i+1];
                    if ($subs == "BC"){
                        $result[$i] = "C";
                        $result[$i+1] = "C";
                        break;
                    }
                }

            }

            break;

        case $randomRule == 5 :
            $lengthInput = strlen($input);

            // getting the rules by looping through the input variable and check it one by one.
            for ($i=0;$i<$lengthInput;$i++) {
                // avoid index out of bound
                if ($i != $lengthInput-1) {
                    $subs = $input[$i] . $input[$i+1];
                    // change AA with A
                    if ($subs == "AA"){
                        $result[$i] = "A";
                        $result[$i+1] = '';
                        break;
                    }
                }

            }

            break;

        case $randomRule == 6 :
            $lengthInput = strlen($input);

            // getting the rules by looping through the input variable and check it one by one.
            for ($i=0;$i<$lengthInput;$i++) {
                // avoid index out of bound
                if ($i != $lengthInput-1) {
                    $subs = $input[$i] . $input[$i+1];
                    if ($subs == "CC"){
                        $result[$i] = "C";
                        $result[$i+1] = '';
                        break;
                    }
                }

            }

            break;
        }

        return implode($result);

    }
}






//echo "print result " . "\n";
//$result = solution2("ABBCC");
////$result = solution("ACC");
//print_r($result . PHP_EOL);
//
//
//$result = solution2("BBCCCAA");
////$result = solution("ACC");
//print_r($result . PHP_EOL);
//
//
//
//$result = solution2("ABCABCABC");
////$result = solution("ACC");
//print_r($result . PHP_EOL);

