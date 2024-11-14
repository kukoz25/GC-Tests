<?php
function create_cycle() {
    $a = new stdClass();
    $b = new stdClass();

    $a->ref = $b;
    $b->ref = $a;

    unset($a);
    unset($b);
}

gc_disable(); // Wyłączenie automatycznego garbage collector

$startTime = microtime(true);
for ($i = 0; $i < 150000; $i++)  {
    create_cycle();
}
$endTime = microtime(true);

echo "Czas wykonania bez garbage collection: " . ($endTime - $startTime) . " sekund\n";

gc_collect_cycles(); // Ręczne uruchomienie garbage collector

$startTime = microtime(true);
for ($i = 0; $i < 150000; $i++)  {
    create_cycle();
}
$endTime = microtime(true);

echo "Czas wykonania z garbage collection: " . ($endTime - $startTime) . " sekund\n";
?>