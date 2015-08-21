	<?php
	$database = 'commons';
	$table = 'testtablewireless';
	$dbhost = 'local';
	$dbuser = 'shawncp';
	$dbpass = '1';
	$conn = mysqli_connect($dbhost, $dbuser, $dbpass);
	mysqli_select_db($conn, $database) or die("Unable to select database");

	?>
