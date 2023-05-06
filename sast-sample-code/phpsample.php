<?php

$conn = new mysqli("127.0.0.1", "user", "MySecretPassword", "db");

function doSqlQuery() {
  $query = "SELECT * FROM users WHERE login = '" . $_POST["login"];

  return mysqli_query($conn, $query);
}

$target = $_GET["target"];

header("Location: " . $target);

echo $_GET["user"];
