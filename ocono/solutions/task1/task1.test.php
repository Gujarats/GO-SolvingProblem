<?php

use PHPUnit\Framework\TestCase;

require 'task1.php';

final class Task1Test extends TestCase{
    protected $inputs = array();

    protected function setup(){
        $this->inputs = array(
            [1,2,3,4],
            [0,0,1,4],
            [2,2,2,4],
            [2,2,2,2],
            [2,2,2,2,1,1,1,1,1,1,2,3,4,2,1,2,1,1],
            [1,2,3,3,3,2,2,21,2,4,2,11],
            [33,33,33,33,1,2,3,3,3,2,2,21,2,4,2,11,33],
        );
    }

    public function test1(){
        $task = new Task1();
        foreach ($this->inputs as $key => $input){
            $result1 = $task->solution($input);
            $result2 = $task->solution2($input);
            if ($result1 != $result2) {
                var_dump("test at index = " . $key);
                print_r($input);
                echo "solution 1 = " . $result1 . PHP_EOL;

                echo "solution 2 = " . $result2;

                $this->assertEquals($result1,$result2);
            }
        }

        $this->assertTrue(true,true);
    }
}
