<?php
function cpu_intensive_task() {
    $result = 0;
    for ($i = 0; $i < 1000000; $i++) {
        $result += sin($i) * tan($i);
    }
    return $result;
}

$load_before = sys_getloadavg();
echo "Obciążenie CPU przed uruchomieniem skryptu: " . implode(", ", $load_before) . "\n";

gc_disable(); // Wyłączenie automatycznego garbage collector

$startCpuTime = microtime(true);
for ($i = 0; $i < 100; $i++) {
    cpu_intensive_task();
}
$endCpuTime = microtime(true);

$load_after = sys_getloadavg();
echo "Obciążenie CPU po uruchomieniu skryptu: " . implode(", ", $load_after) . "\n";
echo "Czas wykonania bez garbage collection: " . ($endCpuTime - $startCpuTime) . " sekund\n";

gc_collect_cycles(); // Ręczne uruchomienie garbage collector

$startCpuTime = microtime(true);
for ($i = 0; $i < 100; $i++) {
    cpu_intensive_task();
}
$endCpuTime = microtime(true);

$load_after_gc = sys_getloadavg();
echo "Obciążenie CPU po uruchomieniu garbage collection: " . implode(", ", $load_after_gc) . "\n";
echo "Czas wykonania z garbage collection: " . ($endCpuTime - $startCpuTime) . " sekund\n";
?>
