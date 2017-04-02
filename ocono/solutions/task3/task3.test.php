<?php

use PHPUnit\Framework\TestCase;

require 'task3.php';

final class Task3Test extends TestCase {

    protected $inputs = array();

    protected function setup(){
        $this->inputs = array(
            "ABBCC",
            "BBCCCAA",
            "ABCCBCBCA",
            "ABCABCABC",
            "ABCCBCBCA",
            "CCCCCCCC",
            "BBBBBBB",
            "AAAAAA",

        );
    }

    // test 1 to test solution funciton using recursive call
    public function test1(){
        $task = new Task3();
        foreach($this->inputs as $key => $input){
            $result = $task->solution($input);
            var_dump($result);
            // check the result if there are no rule anymore
            $rules = $task->getPossibleRules($result);
            $this->assertEmpty($rules);

        }
    }


    // test 1 to test solution2 funciton using while loop 
    public function test2(){
        $task = new Task3();
        foreach($this->inputs as $key => $input){
            $result = $task->solution2($input);
            // check the result if there are no rule anymore
            $rules = $task->getPossibleRules($result);
            $this->assertEmpty($rules);

        }
    }

}
