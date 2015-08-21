<?php



/***************************************************************************************************************************
* Code Maintenance by Jeffrey Dawson from January to April 2014
*
* Additions:	Changes made to the file to allow handling of up to 5 devices per person at a time 
*				student comments removed.
*				style improved. 
*				staff given access to set students wireless device settings. 
*				Links to trouble shooting pages added.
*
* Code Maintenance by Lee Stewart during June and July 2010
*
* Additions: 	Added a function called check_data() that will verify a valid MUNid is entered; Require the user to check
*				a troubleshooting check box if the wireless setup was unSuccessful; Require the user to enter a comment if
*				a client has returned over three times within 30 days and it is a valid return. 
*
*				Added eight troubleshooting check boxes to assist staff if they encounter problems while setting up a laptop
*
*
****************************************************************************************************************************/
	
include ("connection.php"); // contains code to connect to and open the commons database
include 'autoFillArray.php';

echo "

<html>
<head>
<link type=\"text/css\" rel=\"stylesheet\" href=\"wirlessStyle.css\" >

<title>Wireless Admin</title>


<script language=\"javascript\">

/***check_data***************************************************************************
*
* @descript: Checks the editWireless form for errors before when it is submitted
*
* @params: none
*
* @returns: true if there are no errors; false otherwise
*
***************************************************************************************/
function check_data()
{
	if(i==1)
	{
		return false;
	}
	// displays an alert if the MUNid textbox is empty, puts cursor in that textbox,
	// and doesn't submit the form
	if (document.editWireless.MUNid.value == '')
	{
		alert(\"The MUN ID # field is blank\");
		document.editWireless.MUNid.focus();
		return false;
	}

	var txt = document.editWireless.MUNid.value; // a variable to store the value in the MUNid textbox

	// displays an alert if the MUNid textbox doesn't contain 7 or 9 characters (The textbox has a maximum size of 9 characters)
	// puts the cursor in that textbox, and doesn't submit the form.
	if (txt.length < 7 || txt.length == 8)
	{
		alert(\"The MUN ID # field must be between 7 and 9 characters\");
		document.editWireless.MUNid.focus();
		return false;
	}
	
	// if the Setup Successful drop down box option is NO and none of the checkboxes are checked (inner if statement)
	// an alert is displayed, the Restart computer checkbox gets focus, and the form is not submitted.	
	if(document.editWireless.workStatus.options[0].selected && document.editWireless.workStatus.options[0].text == \"No\"
		|| document.editWireless.workStatus.options[1].selected && document.editWireless.workStatus.options[1].text == \"No\"
		|| document.editWireless.workStatus.options[2].selected && document.editWireless.workStatus.options[2].text == \"No\"
		|| document.editWireless.workStatus.options[3].selected && document.editWireless.workStatus.options[3].text == \"No\")
	{	
		if(	!document.editWireless.reboot.checked && 
			!document.editWireless.security.checked && 
			!document.editWireless.locations.checked && 
			!document.editWireless.settings.checked && 
			!document.editWireless.ThirdPartySoftware.checked && 
			!document.editWireless.group.checked && 
			!document.editWireless.durer.checked && 
			!document.editWireless.ipv6.checked)
			
		{
				alert(\"You must select a checkbox\");
				document.editWireless.reboot.focus();
				return false;
		}
	}
	
	// The outer if statement gets the value in the hidden values checkbox. If the client has returned before,
	// the inner if statement will display an alert if there is no comment in the Staff Comments field, put the
	// cursor in that field, and not submit the form. 
	if (eval(document.editWireless.visits.value) > 1)
	{
		if (document.editWireless.returned.options[1].selected && document.editWireless.staffcomments.value=='')
		{
				alert(\"You must enter a comment in the Staff Comments field\");
				document.editWireless.staffcomments.focus();
				return false;	
		}
	}
	return true;	// if there are no errors, true is returned and the form will submit.
}

function myFunction(){
var i=1;
return false;
}
</script>
</head>

<body bgcolor=\"#488fcd\">
<center><div id=logo></center>
<h1 align=\"center\">Wireless Log Database Search/Edit</h1>";	

				
function BrandName($brand)
{
	//get staff initials from database
	echo "<option>$brand</option>";	
	
	$brand = autobrand($section);
	$auto_fill = "edit";
	foreach ($brand as $value) 
	{
		echo "<option>" . $value . "</option>";
	}
}

function OSName($os)
{
	//get staff initials from database
	echo "<option>$os</option>";	
	
	$os = autoos($section);
	$auto_fill = "edit";
	foreach ($os as $value) 
	{
		echo "<option>" . $value . "</option>";
	}
}
function fillTable($OSt, $statusA, $GETbrand, $os)
			{
				$OSType=$_GET[$OSt];
				$brand  = $_GET[$GETbrand];
				if($status= @$_GET[$statusA]){}
				else
				{
					$status = @$_GET['status'];
				}
				echo"
				<td><input type=\"text\" size=\"7\" name=\"OperatingSystemType1\" value=\"" . $OSType . "\"></td>
				<td>
					<select name=\"brand\" id=\"brand\">";
					BrandName($brand);
					echo "</select>
				</td>	
				<td>
					<select name=\"os\" id=\"os\">";
					OSName($os);
					echo "</select>
				</td>
				<td><select name=\"workStatus1\" id=\"workStatus1\">";
				setupStatus($status);
			}
// gets values from the modifyWireless.php file 
	$id = $_GET['id'];
	$date = $_GET['date'];
	$firstName = $_GET['firstName'];
	$lastName = $_GET['lastName'];
	$MUNid = $_GET['MUNid'];
	$OSType=$_GET['OperatingSystemType'];
	$brand  = $_GET['brand'];
	$os = $_GET['os'];
	@$os2 = $_GET['os2'];
	@$os3 = $_GET['os3'];
	@$os4 = $_GET['os4'];
	@$os5 = $_GET['os5'];
	$clientType = $_GET['clientType'];
	$staff = $_GET['staff'];
	@$status = $_GET['status'];
	$staffcomments = $_GET['staffcomments'];
	$count = $_GET['NUMdevices'];
	$section = $_GET['section'];
	
function setupStatus($status)
{
			

			/******** Section added by Lee Stewart **********/
			//set up secessull code
			if ($status == '')  // for clients that just entered a record the status will be null. 
			{					// This code will populate the drop down box with the required options.
				echo "<option value=\"Yes\">Yes</option>";
				echo "<option value=\"No\">No</option>";
				echo "<option value=\"Pending\">Pending</option>";
				echo "<option value=\"Left\">Left</option>";
				echo "<option value=\"abandoned\">abandoned</option>";
				echo "<option value=\"escalated\">sent to C&C</option>";
			}
			else	// For existing records, the status will have a value. 
			{		
				echo "<option value=$status>$status</option>"; // and that status will be the first option in the drop down list
				//The code in the blocks below correct an issue where double entries were appearing in the drop down list. 
				if ($status == "Yes")
				{
					echo "<option value=\"No\">No</option>";
					echo "<option value=\"Pending\">Pending</option>";
					echo "<option value=\"Left\">Left</option>";
					echo "<option value=\"abandoned\">abandoned</option>";
					echo "<option value=\"escalated\">sent to C&C</option>";
				}
				elseif ($status == "No")
				{
					echo "<option value=\"Yes\">Yes</option>";
					echo "<option value=\"Pending\">Pending</option>";
					echo "<option value=\"Left\">Left</option>";
					echo "<option value=\"abandoned\">abandoned</option>";
					echo "<option value=\"escalated\">sent to C&C</option>";
				}
				elseif ($status == "Pending")
				{
					echo "<option value=\"Yes\">Yes</option>";
					echo "<option value=\"No\">No</option>";
					echo "<option value=\"Left\">Left</option>";
					echo "<option value=\"abandoned\">abandoned</option>";
					echo "<option value=\"escalated\">sent to C&C</option>";
				}
				elseif($status == "abandoned"){
					echo "<option value=\"Yes\">Yes</option>";
					echo "<option value=\"Pending\">Pending</option>";
					echo "<option value=\"No\">No</option>";
					echo "<option value=\"Left\">Left</option>";
					echo "<option value=\"escalated\">sent to C&C</option>";
				}
				elseif($status == "escalated"){
					echo "<option value=\"Yes\">Yes</option>";
					echo "<option value=\"Pending\">Pending</option>";
					echo "<option value=\"No\">No</option>";
					echo "<option value=\"Left\">Left</option>";
					echo "<option value=\"abandoned\">abandoned</option>";
				}
				else
				{
					echo "<option value=\"Yes\">Yes</option>";
					echo "<option value=\"No\">No</option>";
					echo "<option value=\"Pending\">Pending</option>";
					echo "<option value=\"abandoned\">abandoned</option>";
				}
			}
}
	
	
	
	
	if($staff == NULL){
		$staff = "N/A";
	}

	if (preg_match("/\bUndergraduate\b/i", $clientType)){  //jr: replaced eregi with preg_match
		$clientType = "Undergrad";
	}



// display the selected log
	echo "<p>You selected &quot;" . $firstName . " " . $lastName . "&quot;<br></p>";

echo "<table id=\"table1\" width=\"100%\" border=\"1\">
	<tr align = \"center\">
	<td style=\"display:none;\" ><b>ENTRY ID</b></td>
	<td><b>Date/Time<br>Entered</b></td>
	<td><b>First Name</b></td>
	<td><b>Last Name</b></td>
	<td><b>MUN ID#</b></td>
		<td><b>Student?<br>Employee?</b></td>
	<td><b>Staff</b></td>
	<td><b>Operating<br>System Type</b></td>
	<td><b>Laptop<br>Brand</b></td>
	<td><b>Operating<br>System</b></td>
	<td><b>Setup<br>Successful?</b></td>	
	<td style=\"display:none;\"><b>RETURNED</b></td>
	<td><b>Staff Comments</b></td>
	<td style=\"display:none;\" ><b>count</b></td>
	</tr>";


	if($count==NULL){$count=1;}
$counts1=$count;
$counts2=$count;
$i=1;





/***code for top table************************************
*this code loops based on the counter passed in by the MySQLtable
*it out puts the information for the devices that where submitted at the same time. 
********************************************************/

for(; $count>0; $count--)
	{	
		$staff = $_GET['staff'];
		if($staff == NULL)
		{
			$staff = "N/A";
		}

		if (preg_match("/\bUndergraduate\b/i", $clientType))  //jr: replaced eregi with preg_match
		{
			$clientType = "Undergrad";
		}
		if($i!=1)
		{
			$id++;
		}

		echo "<form name=\"editWireless\" action=\"editWirelessLog.php\" method=\"GET\" onSubmit=\"return check_data()\">

			
			<input type=\"hidden\" name=\"id\" value=\"" . $id . "\">
			<input type=\"hidden\" name=\"date\" value=\"" . $date . "\">
			<tr align = \"center\">
			";if($i==1)
			{
				$i--;
				echo"
				<td style=\"display:none;\">" . $id . "</td>
				<td>" . $date . "</td>
				<td><input type=\"text\" size=\"7\" name=\"FirstName\" value=\"" . $firstName . "\"></td>
				<td><input type=\"text\" size=\"7\" name=\"LastName\" value=\"" . $lastName . "\"></td>
				<td><input type=\"text\" size=\"7\" maxlength=9 name=\"MUNid\" value=\"" . $MUNid . "\"></td>
														
				<td><select name=\"stuEmp\">
				<option>$clientType</option>
				<option value=\"Undergraduate\">Undergrad</option>
				<option value=\"Alumni\">Alumni</option>
				<option value=\"Graduate\">Graduate</option>
				<option value=\"Faculty\">Faculty</option>
				<option value=\"Staff\">Staff</option>
				<option value=\"Other\">Other</option>
				</select>
				</td>			
				<td>
					<select name=\"staff\" id=\"staff\">";
					//check to see if a staff member is assigned to this client yet
					if ($staff == "N/A") 
					{
						$helped_already = FALSE;
					}
					else 
					{
						$helped_already = TRUE;
					}
					
					
					//get staff initials from database
					
					echo "<option>$staff</option>";	
						
						$staff = autostaff($section);
						$auto_fill = "edit";
						foreach ($staff as $value) 
						{
							echo "<option>" . $value . "</option>";
						}
					echo "</select>";
					}
					else
					{
						echo"
						<td style=\"display:none;\" >" . $id . "</td>
						<td>" . $date . "</td>
						<td> </td>
						<td> </td>
						<td> </td>
						<td> </td>
						<td> </td>";
					}

			if($counts1==1)
			{
				$OSType1=$_GET['OperatingSystemType'];
				$brand  = $_GET['brand'];

				if($status1 = @$_GET['status1']){}
				else
				{
					$status1 = @$_GET['status'];
				}
				echo"
				<td><input type=\"text\" size=\"7\" name=\"OperatingSystemType1\" value=\"" . $OSType1 . "\"></td>
				<td>
					<select name=\"brand\" id=\"brand\">";
					BrandName($brand);
					echo "</select>
				</td>	
				<td>
					<select name=\"os\" id=\"os\">";
					OSName($os);
					echo "</select>
				</td>";
				echo "<td><select name=\"workStatus1\" id=\"workStatus1\">";
				setupStatus($status1);
				echo 	"</tb><td><textarea  id=\"text1\"   rows=\"5\" cols=\"20\" name=\"staffcomments\"></textarea>
				</td>";
			}
			if($counts1==2)
			{
				$OSType2=$_GET['OperatingSystemType2'];
				$brand2  = $_GET['brand2'];

				if($status2 = @$_GET['status2']){}
				else
				{ 
					$status2 = @$_GET['status'];
				}

				$counts1--;
				echo"
				<td><input type=\"text\" size=\"7\" name=\"OperatingSystemType2\" value=\"" . $OSType2 . "\"></td>

				<td>
					<select name=\"brand2\" id=\"brand2\">";
					BrandName($brand2);
					echo "</select>
				</td>	
				<td>
					<select name=\"os2\" id=\"os2\">";
					OSName($os2);
					echo "</select>
				</td>	";
				echo "<td><select name=\"workStatus2\" id=\"workStatus2\">";
				setupStatus($status2);
				echo 	"</tb><td><textarea  id=\"text1\"   rows=\"5\" cols=\"20\" name=\"staffcomments2\"></textarea>
				</td>";
			}
			if($counts1==3)
			{
				
				$OSType3=$_GET['OperatingSystemType3'];
				$brand3  = $_GET['brand3'];

				if($status3 = @$_GET['status3']){}
				else
				{ 
					$status3 = @$_GET['status'];
				}
				$counts1--;
				echo"
				<td><input type=\"text\" size=\"7\" name=\"OperatingSystemType3\" value=\"" . $OSType3 . "\"></td>

				<td>
					<select name=\"brand3\" id=\"brand3\">";
					BrandName($brand3);
					echo "</select>
				</td>	
				<td>
					<select name=\"os3\" id=\"os3\">";
					OSName($os3);
					echo "</select>
				</td>	";
				echo "<td><select name=\"workStatus3\" id=\"workStatus3\">";
				setupStatus($status3); 
				echo 	"</tb><td><textarea  id=\"text1\"   rows=\"5\" cols=\"20\" name=\"staffcomments3\"></textarea>
				</td>";				
			}
			if($counts1==4)
			{
				$OSType4=$_GET['OperatingSystemType4'];
				$brand4  = $_GET['brand4'];

				if($status4 = @$_GET['status4']){}
				else
				{ 
					$status4 = @$_GET['status'];
				}
				$counts1--;
				echo"
				<td><input type=\"text\" size=\"7\" name=\"OperatingSystemType4\" value=\"" . $OSType4 . "\"></td>
				
				<td>
					<select name=\"brand4\" id=\"brand4\">";
					BrandName($brand4);
					echo "</select>
				</td>	
				<td>
					<select name=\"os4\" id=\"os4\">";
					OSName($os4);
					echo "</select>
				</td>	";
				echo "<td><select name=\"workStatus4\" id=\"workStatus4\">";
				setupStatus($status4);
				echo 	"</tb><td><textarea  id=\"text1\"   rows=\"5\" cols=\"20\" name=\"staffcomments4\"></textarea>
				</td>";				
			}
			if($counts1==5)
			{
				$OSType5=$_GET['OperatingSystemType5'];
				$brand5  = $_GET['brand5'];
				if($status5 = @$_GET['status5']){}
				else
				{ 
					$status5 = @$_GET['status'];
				}
				$counts1--;
				echo"
				<td><input type=\"text\" size=\"7\" name=\"OperatingSystemType5\" value=\"" . $OSType5 . "\"></td>
	
				<td>
					<select name=\"brand5\" id=\"brand5\">";
					BrandName($brand5);
					echo "</select>
				</td>	
				<td>
					<select name=\"os5\" id=\"os5\">";
					OSName($os5);
					echo "</select>
				</td>	";
				echo "<td><select name=\"workStatus5\" id=\"workStatus5\">";
				setupStatus($status5);
				echo 	"<td><textarea  id=\"text1\"   rows=\"5\" cols=\"20\" name=\"staffcomments5\"></textarea>
				</td>";				
			}
				
					
			echo"</td>";

			
			/******end of this section *******************************************/
			echo "</select>";

			/******** Section added by Lee Stewart on June 16, 2010 switch block added by Jeffrey Dawson April 2014 to account for more than one device************/
			// This section queries the wireless table to see if a client has returned within 30 days and creates a drop down box 
			// with different option depending on the result
			// These three statements are required to generate the query and get the results.
			switch($counts2)
			{
			
				case 1: 
					$visitQuery = 	"SELECT count(*) 
							FROM $database.$table
							WHERE dtDate > DATE_SUB(now(), INTERVAL 30 DAY)
							AND txtMUNID = '$MUNid' AND (txtOS = '$os') 
							GROUP BY txtOS
							ORDER BY count(*) DESC
							LIMIT 1";
				break;
			
				case 2: 
					$visitQuery = 	"SELECT count(*) 
							FROM $database.$table
							WHERE dtDate > DATE_SUB(now(), INTERVAL 30 DAY)
							AND txtMUNID = '$MUNid' AND (txtOS = '$os' OR txtOS = '$os2') 
							GROUP BY txtOS
							ORDER BY count(*) DESC
							LIMIT 1";
				break;
				case 3: 
					$visitQuery = 	"SELECT count(*) 
							FROM $database.$table
							WHERE dtDate > DATE_SUB(now(), INTERVAL 30 DAY)
							AND txtMUNID = '$MUNid' AND (txtOS = '$os' OR txtOS = '$os2' OR txtOS = '$os3') 
							GROUP BY txtOS
							ORDER BY count(*) DESC
							LIMIT 1";
				break;
				case 4: 
					$visitQuery = 	"SELECT count(*) 
							FROM $database.$table
							WHERE dtDate > DATE_SUB(now(), INTERVAL 30 DAY)
							AND txtMUNID = '$MUNid' AND (txtOS = '$os' OR txtOS = '$os2' OR txtOS = '$os3' OR txtOS = '$os4') 
							GROUP BY txtOS
							ORDER BY count(*) DESC
							LIMIT 1";
				break;
				case 5: 
					$visitQuery = 	"SELECT txtOS, count(*) 
							FROM $database.$table
							WHERE dtDate > DATE_SUB(now(), INTERVAL 30 DAY)
							AND txtMUNID = '$MUNid' AND (txtOS = '$os' OR txtOS = '$os2' OR txtOS = '$os3' OR txtOS = '$os4' OR txtOS = '$os5') 
							GROUP BY txtOS
							ORDER BY count(*) DESC
							LIMIT 1";
				break;
			}

			$visitResults = 	mysqli_query($conn, $visitQuery);
			$visits = mysqli_fetch_array($visitResults);
			$visits = $visits[0];	
			
			// Creates a dynamic drop down box
			echo "<td style=\"display:none;\"> <select name = \"returned\">";
				if ($visits > 3)	//if the client has returned within 30 days the drop down contains yes and no with yes selected
				{ 
					echo "<option value=\"no\">No</option>";
					echo "<option value=\"yes\" selected=\"selected\">Yes</option>";	
				}
				else	// if client hasn't returned within 30 days, only no is displayed
				{
					echo "<option value=\"no\">No</option>";	
				}
				echo "</select></td>";

			/********** end of section *****************************************/	

				echo "<td style=\"display:none;\" > <input type=\"text\" size=\"1\" name=\"count\" value=\"" . $counts2 . "\"></td>
				<input type=\"hidden\" name=\"helped\" value=\"" . $helped_already . "\" /> <!-- hidden box to tell if the peon is being helped -->
				<br>";
				//taking the delete button out till I can 
				//find out how to create a warning message 
				//echo "<input align=\"centre\" type=\"submit\" name=\"delete\" value=\"Delete\">";
				//~Lee
		}

	echo "</td>
		</tr>
			</table>";
			
		


			
			echo "<div align=\"center\"><input class=\"Submitbutton\"  type=\"submit\" name=\"edit\" value=\"Update All\"> </div>";// put in by jeffrey






	/**** Section added by Lee Stewart on June 14, 2010 **check boxes******************/
	if ($visits > 3) // for repeat clients a red warning message is displayed indicating the number of times they returned.
	{
		echo "<h2 align=center><font color=yellow> Repeat Client. They have returned with " . $visits . " device problems in the last 30 days</font></h2>";
	}
	
	// This section of code sets up the check boxes that will be used for troubleshooting.
	// They are contained within a table with no borders
	echo "<center><table border = 1 cellspacing=0 cellpadding=2>
	<caption align = \"top\"><h3>Troubleshooting Steps</h3><center><a href=http://www.cs.mun.ca/~thecommons/wireless/ target=\"_blank\"><font color=\"fefefe\">Wireless Guide</font> </a></caption>
	<caption align = \"top\"><center><a href=https://cgtools.munet.mun.ca/durer/ target=\"_blank\"><font color=\"fefefe\">Durer2.0</font> </a></caption>
	<caption align = \"top\"><center><a href=https://www.mun.ca/cc/services/network/wireless/munlogin/acceptable_use.php target=\"_blank\"><font color=\"fefefe\">DAT form</font> </a></caption>
	<caption align = \"top\"><center><a href=https://www.mun.ca/cc/services/network/wireless/munlogin/acceptable_use.php?guest target=\"_blank\"><font color=\"fefefe\">GUEST DAT form</font> </a></caption>

	<tr>
		  
		   <td valign=\"top\">
			<input type=\"checkbox\" name=\"reboot\" value=\"reboot laptop\" />
		   </td>
		   <td valign=\"top\">
			Restart computer
		   </td>
		   <td valign=\"top\">
			<input type=\"checkbox\" name=\"2nd device\" value=\"2nddevice\" />
		   </td>
		   <td valign=\"top\">
			2nd Wireless Device
		   </td>   
		</tr>
		<tr>
		   <td valign=\"top\">
			<input type=\"checkbox\" name=\"settings\" value=\"settings are correct\" />
		   </td>
		   <td valign=\"top\">
			Check wireless settings
		   </td>
		   <td valign=\"top\">
			<input type=\"checkbox\" name=\"group\" value=\"others affected\" />
		   </td>
		   <td valign=\"top\">
			Are others affected?
		   </td>   		  
		</tr>
		<tr>
		   <td valign=\"top\">
			<input type=\"checkbox\" name=\"durer\"  value=\"password correct in durer\" />
		   </td>
		   <td valign=\"top\">
			Check password in durer
		   </td>
		    <td valign=\"top\">
			<input type=\"checkbox\" name=\"locations\" value=\"other\" />
		   </td>
		   <td valign=\"top\">
			 other
		   </td>
		</tr>
		</table></center>";
		
		
		
		
		 echo "<input type='hidden' name='section' value='$section'>";
	// THIS HIDDEN FIELD IS VITAL TO THIS PROGRAM AS IT ENABLES JAVASCRIPT AND PHP TO WORK IN HARMONY.
	// IT IS CRITICAL TO THE FUNCTIONING OF THE check_data FUNCTION BETWEEN THE SCRIPT TAGS AT THE TOP
	echo "<input type=\"hidden\" name=\"visits\" value= \"" . $visits . "\"/>"; 
	/********** end of section *********************************************/
	echo "</form>"; // THE END OF THE editWireless FORM
	
		echo "<form method=\"POST\" action=\"searchWirelessLessinfo.php?Search=&tableField=vchWorkStatus&section=2&SW1=Search\">
	 <button class=\"myButton\" type=\"submit\">Return</button><table>
	</table></form>";
	
	
	
	// rows to return
	$limit=20; 
	
	
	
	
	//not moved up higher incode
	//query to find all other times this person has come for wireless
	$query = "SELECT * FROM $database.$table WHERE txtMUNID = '$MUNid' ORDER BY iID DESC";
	$result = mysqli_query($conn, $query) or die("Couldn't execute query");
	$numrows=mysqli_num_rows($result);
	$result1=$result;
	

	
// If we have no results.
	if ($numrows == 0)
	{
		echo "<h4>Results</h4>
		<p>Sorry, your search in " . $fields . " for: &quot;" . $trimmed . "&quot; returned zero results</p>";

		//Link to search page
	
		echo "<p>Return to search pages.</p>
		<form method=\"Post\" action=\"searchWirelessLessinfo.php?Search=&tableField=vchWorkStatus&section=2&SW1=Search\"><table>
		<td><input class=\"Submitbutton\" id=\"returnButton\" type=\"submit\" value=\"Return\"></td>
		</table>
		</form>	
		</body>
		</html>";
	}


// next determine if s has been passed to script, if not use 0
	if (empty($s)){
		$s=0;
		}

// begin to show results set
	echo "Results";
	$count = 1 + $s ;
/********code for the header of the bottom table (past results )******************/
	echo "<table width=\"100%\" border=\"1\">
	<tr align =\"center\">
	<td><b>Entry ID</b></td>
	<td><b>Date/Time<br>Entered</b></td>
	<td><b>First Name</b></td>
	<td><b>Last Name</b></td>
	<td><b>MUN ID#</b></td>
	<td><b>Operating<br>System Type</b></td>
	<td><b>Laptop<br>Brand</b></td>
	<td><b>Operating<br>System</b></td>
	<td><b>Student?<br>Employee?</b></td>
	<td><b>Staff</b></td>
	<!--<td><b>RETURNED</b></td>-->
	<td><b>SETUP<br>Successful?</b></td>
	<!--td><b>COMMENTS</b></td-->
	<td><b>Staff Comments</b></td>
	
	</tr>";

	
	
/********code for the bottom table (past results )******************/

while (($row= mysqli_fetch_array($result1))){ 

		echo "<form  name=\"modifyWireless\" method=\"GET\" action=\"modifyWirelessLog.php\" onsubmit=\"return\">
		<input type=\"hidden\" name=\"id\" value=\"" . $row['iID'] . "\">
		<input type=\"hidden\" name=\"date\" value=\"" . $row['dtDate'] . "\">
		<input type=\"hidden\" name=\"firstName\" value=\"" . $row['txtFirstName'] . "\">
		<input type=\"hidden\" name=\"lastName\" value=\"" . $row['txtLastName'] . "\">
		<input type=\"hidden\" name=\"MUNid\" value=\"" . $row['txtMUNID'] . "\">
		<input type=\"hidden\" name=\"OperatingSystemType\" value=\"" . $row['OSType'] . "\">
		<input type=\"hidden\" name=\"brand\" value=\"" . $row['txtBrand'] . "\">
		<input type=\"hidden\" name=\"os\" value=\"" . $row['txtOS'] . "\">
		<input type=\"hidden\" name=\"clientType\" value=\"" . $row['vchClientType'] . "\">
		<input type=\"hidden\" name=\"staff\" value=\"" . $row['txtStaffInitial'] . "\">
		<input type=\"hidden\" name=\"status\" value=\"" . $row['vchWorkStatus'] . "\">
		
		<input type=\"hidden\" name=\"staffcomments\" value=\"" . $row['txtStaffComments'] . "\">
		<tr align =\"center\">
		<td align = \"center\">" . $row['iID'] . "</td>
		<td>" . $row['dtDate'] . "</td>
		<td>" . $row['txtFirstName'] . "</textarea></td>
		<td>" . $row['txtLastName'] . "</textarea></td>
		<td>" . $row['txtMUNID'] . "</textarea></td>
		<td>" . $row['OSType'] . "</textarea></td>
		<td>" . $row['txtBrand'] . "</textarea></td>
		<td>" . $row['txtOS'] . "</textarea></td>
		<td>" . $row['vchClientType'] . "</textarea></td>";
		if($row['txtStaffInitial'] == NULL){
			echo "<td>N/A</textarea></td>";
		}
		else{echo "<td>" . $row['txtStaffInitial'] . "</textarea></td>";}
		echo "<td>" . $row['vchWorkStatus'] . "</textarea></td>";
		
		if($row['txtStaffComments'] == NULL){
			echo "<td>N/A</textarea></td>";
		}
		else{echo "<td>" . $row['txtStaffComments'] . "</textarea></td>";}
		//echo "<td><input class='Submitbutton' type='submit' value='Edit'></td>
		echo "
		</tr>
		</form>";

}
	
	
	mysqli_close($conn);
?>
<br>
</body>
</html>
