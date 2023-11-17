<?php

$_ROOT = __DIR__ . "/../";

/**
 * requires a file from $_ROOT directory
 */
function requireFromRoot(string $require) {
    global $_ROOT;
    require_once $_ROOT . $require;
}