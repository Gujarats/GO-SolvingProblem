<?php

use PHPUnit\Framework\TestCase;

require 'app.php';

final class EmailTest extends TestCase
{
    public function test1(){
        $trying = new Trying();
        //var_dump($trying->getHello());
        $this->assertEquals("helloworld",$trying->getHello());
    }
}
?>
