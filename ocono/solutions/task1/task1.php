<?php

class Task1 {
    public function solution(&$A) {
        $N = sizeof($A);
        $result = 0;
        for ($i = 0; $i < $N; $i++)
            for ($j = 0; $j < $N; $j++)
                // checking the same number value
                if ($A[$i] == $A[$j]){
                    $result = max($result, abs($i - $j));

                }

        return $result;
    }

    public function solution2($A)  {
        $N = sizeof($A);


        // create a storgae and save as key value.
        // key = value of the number
        // value = {
        //  repetation : 12,
        //  minIndex : 1,
        //  maxIndex: 2
        // }
        $storage = array();
        $lastNumber = 0;
        $result = 0;

        $minIndex = 0;
        $maxIndex = 0;


        if ($A[0] == $A[$N-1]){
            return $N-1;
        }else{

            // store the new value as an array object
            $storage[$A[0]] = array(
                "minIndex" => 0,
                "maxIndex" => 0,
                "repetation" => 1,
            );

            // assign to last number
            $lastNumber = $A[0];
            // begin looping finding the max total same number
            for ($i = 1; $i < $N; $i++){
                // compare lastNumber to the current index value
                if ($lastNumber == $A[$i]){
                    // assign current index to $result. max index the same number.
                    $storage[$A[$i]]["maxIndex"] = $i;
                    $storage[$A[$i]]["repetation"] = $storage[$A[$i]]["repetation"] +1;

                }else{
                    $lastNumber = $A[$i];
                    // check if new number has previous storage
                    if (!empty($storage[$A[$i]])){
                        // assign current index to $result. max index the same number.
                        $storage[$A[$i]]["maxIndex"] = $i;
                        $storage[$A[$i]]["repetation"] = $storage[$A[$i]]["repetation"] +1;
                    }else{
                        // store new number
                        $storage[$A[$i]] = array(
                            "minIndex" => $i,
                            "maxIndex" => $i,
                            "repetation" => 1,
                        );

                    }
                }
            }

            // find the max repetation and get minIndex and maxIndex
            $result = 0;
            foreach($storage as $key => $object) {
                $difference  = abs($object["minIndex"] - $object["maxIndex"]);
                if ($result< $difference){
                    $result = $difference;
                }
            }

            return $result;

        }
    }
}

