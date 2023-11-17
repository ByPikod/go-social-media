<?php

$db = null;

function pgConnect()
{
    global $_CONF;
    if (!extension_loaded('pdo_pgsql')) {
        die('PGSQL extension not loaded');
    }
    $query = sprintf(
        "pgsql:host=%s port=%s user=%s password=%s dbname=%s",
        $_CONF["DB"]["HOST"],
        "5432",
        $_CONF["DB"]["USER"],
        $_CONF["DB"]["PASS"],
        $_CONF["DB"]["DB"]
    );
    $db = new PDO($query, null, null, [
        PDO::ATTR_ERRMODE => PDO::ERRMODE_EXCEPTION,
        PDO::ATTR_CASE => PDO::CASE_NATURAL,
        PDO::ATTR_DEFAULT_FETCH_MODE => PDO::FETCH_ASSOC,
    ]);
    return $db;
}

$db = pgConnect();
