<?php
declare(strict_types=1);

class Trying
{
    public $storage = array();

    public function getHello():string
    {
        return "helloworlds";
    }

    public function getValue($key,$a,$b,$c):array
    {
        $storage[$key] = array(
            "a" => $a,
            "b" => $b,
            "c" => $c,
            1 => $c,
        );
        return $storage;
    }

    public function change($input){
        $input  = $input . " cekcceok";
    }
}


$test = new Trying();
$input = "lolo";

$test->change($input);

echo $input . PHP_EOL;
