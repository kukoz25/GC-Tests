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

$startMemory = memory_get_usage();
for ($i = 0; $i < 10000; $i++) {
    create_cycle();
}
$endMemory = memory_get_usage();

echo "Zużycie pamięci bez garbage collection: " . ($endMemory - $startMemory) . " bitów\n";

gc_collect_cycles(); // Ręczne uruchomienie garbage collector

$startMemory = memory_get_usage();
for ($i = 0; $i < 10000; $i++) {
    create_cycle();
}
$endMemory = memory_get_usage();

echo "Zużycie pamięci z garbage collection: " . ($endMemory - $startMemory) . " bitów\n";
?>
