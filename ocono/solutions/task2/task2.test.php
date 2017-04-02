<?php

declare(strict_types=1);

use PHPUnit\Framework\TestCase;

require 'task2.php';

// create test case class
class TestObject {
    private $input;
    private $expected;

    public function __construct($input,$expected){
        $this->input = $input;
        $this->expected = $expected;
    }

    public function getInput(){
        return $this->input;
    }

    public function getExpected(){
        return $this->expected;
    }
}

final class Task2Test extends TestCase {

    
    protected $inputs = array();
    protected $task;

    protected function setup(){
        $this->task = new Task2();

        $this->inputs = array(
            new TestObject("13 12",12),
            new TestObject("13 DUP",13),
            new TestObject("13 DUP DUP DUP",13),
            new TestObject("DUP DUP DUP",-1),
            new TestObject("+",-1),
            new TestObject("-",-1),
            new TestObject("POP",-1),
            new TestObject("asdfasdf",-1),
            new TestObject("19000000",-1),
            new TestObject("13 POP DUP",-1),
            new TestObject("13 DUP 4 POP 5 DUP + DUP + -",7),
            new TestObject("13 DUP 4 POP 5 DUP + DUP + - 19000000",-1),
            new TestObject("13 DUP 4 HEH 5 DUP + DUP + - 19000000",-1),
            new TestObject("13DUP 4 HEH 5 DUP + DUP + - 19000000",-1),
            new TestObject("5 6 + -",-1),
            new TestObject("3 DUP 5 - -",-1),
        );
    }

    public function testOperationNumber(){
        
        // define the arguments
        $stack = new Stack(12,null,null);
        $nextStack= new Stack(1,$stack,null);
        $stack->next = $nextStack;

        $expected = 4;
        
        $result = $this->task->operation($nextStack,4);
        
        $this->assertEquals($result["error"],0);
        $this->assertEquals($result['stack']->value,4);

    }

    public function testOperationDUP(){
        $stack = new Stack(12,null,null);
        $nextStack= new Stack(1,$stack,null);
        $stack->next = $nextStack;

        $result = $this->task->operation($nextStack,"DUP");
        // the result must not have an error value -1
        $this->assertEquals($result["error"],0);

        // check if we duplication the top stack
        $resultStack = $result["stack"];
        $this->assertEquals($resultStack->value ,$resultStack->prev->value);

    }

    public function testOperationPOP(){
        $stack = new Stack(12,null,null);
        $nextStack= new Stack(1,$stack,null);
        $stack->next = $nextStack;


        $result = $this->task->operation($nextStack,"POP");
        // the result must not have an error value -1
        $this->assertEquals($result["error"],0);

        // check if we duplication the top stack
        $resultStack = $result["stack"];
        $this->assertEquals($resultStack->value ,12);
    }

    public function testOperationPlus(){
        $stack = new Stack(12,null,null);
        $nextStack= new Stack(1,$stack,null);
        $stack->next = $nextStack;


        $result = $this->task->operation($nextStack,"+");
        // the result must not have an error value -1
        $this->assertEquals($result["error"],0);

        // check if we duplication the top stack
        $resultStack = $result["stack"];
        $this->assertEquals($resultStack->value ,13);
    }

    public function testOperationMinus(){
        $stack = new Stack(5,null,null);
        $nextStack= new Stack(7,$stack,null);
        $stack->next = $nextStack;
        
        $expected = 2;

        $result = $this->task->operation($nextStack,"-");
        // the result must not have an error value -1
        $this->assertEquals($result["error"],0);

        // check if we duplication the top stack
        $resultStack = $result["stack"];
        $this->assertEquals($resultStack->value ,$expected);
    }


    public function testSolution(){

        foreach($this->inputs as $key => $testCase){
            $input = $testCase->getInput();
            $expected = $testCase->getExpected();
            $result = $this->task->solution($input);
            if ($result != $expected){
                var_dump("test index at = " . $key);
                var_dump($input);
                $this->assertEquals($expected,$result);
            }
        }

        $this->assertTrue(true);
    }
}

