<?php

declare(strict_types=1);

class Stack {
    public $value;
    public $prev;
    public $next;
    
    function __construct($value,$prev,$next)
    {
        $this->value = $value;
        $this->prev= $prev;
        $this->next= $next;
    }
}


class Task2 {

    public function operation($stack,$command):array{
        $result = array(
            "stack"=> $stack,
            "error"=> 0,
        );

        

        switch($command){
        case "DUP" :
            //checking if stack has a value
            if($stack==null){
                $result["error"] = -1;
            }else{
                if($stack->value == null){
                    $result["error"] = -1;
                }else{
                    $newStack = new Stack($stack->value,$stack,null);
                    //$stack->next = $newStack;
                    //$newStack->prev = $stack;
                    $result["stack"] = $newStack;
                }

            }
            break;
        case "POP" :
            if($stack->value == null){
                $result["error"] = -1;
            }else{
                $newStack = $stack->prev;
                $stack = null;
                $result["stack"] = $newStack;
            }
            break;
        case "+" :
            // check if stack has at leat two stack value
            if($stack->value == null || $stack->prev == null){
                $result["error"] = -1;
            }else{
                $resultPlus = $stack->value + $stack->prev->value; 
                $newStack = $stack->prev;
                $newStack->value = $resultPlus;
                $result["stack"] = $newStack;
            }
            break;
        case "-" :
            // check if stack has at leat two stack value
            if($stack->value == null || $stack->prev == null){
                $result["error"] = -1;
            }else{
                $resultMinus = $stack->value - $stack->prev->value; 
                if($resultMinus < 0 ){
                    $result["error"] = -1;
                }else{
                    $newStack = $stack->prev;
                    $newStack->value = $resultMinus;
                    $result["stack"] = $newStack;
                }
            }

            break;
        case is_int($command):
            //check integer value for validation
            // only accept 0 - 2^20-1
            if($command >1048575 ){
                $result["error"] = -1;
            }else{
                // add new number to the stack
                if (empty($stack->value)){
                    $result["stack"] = new Stack($command,null,null);
                }else{
                    $newStack = new Stack($command,$stack,null);
                    $stack->next = $newStack;
                    $result["stack"] = $newStack;
                }

            }

            break;
        default :
            // no comment match
            $result["error"] = -1;
        }

        return $result;
    }

    public function solution($operation):int{
        $inputs = explode(" ",$operation);

        // converting the int value to int from string
        foreach($inputs as $key => $value){
            if ($value != "DUP" && $value != "POP" && $value != "+" && $value != "-"){
                // $value is int
                $inputs[$key] = intval($value);
            }
        }

        // return -1 if the first inputs is not int 
        if (!is_int($inputs[0])){
            return -1;
        }

        // create empty stack
        $stack = new Stack(null,null,null);
        // add all inputs to operation
        $result = array(
            "stack" => $stack,
            "error" => 0,
        );

        // repeating insert inputs to operation .
        foreach($inputs as $key => $command){
            $result = $this->operation($stack,$command);
            if ($result["error"] == -1){
                return -1;
            }else{

                $stack = $result["stack"];
            }
        }

        // check if we got error from the operation function.
        if($result["error"]!=-1){
            return $result["stack"]->value;
        }else{
            return -1;
        }

    }
}



