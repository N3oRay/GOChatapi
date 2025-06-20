

<form action="index.php" method="post">
 <p>Bienvenu :</p>
 <p>Chat : <input type="text" name="nom" placeholder="Ex : Bonjour" autofocus /></p>
 
 <p>
       Cochez un choix pour ton donner votre avis :<br />
       <input type="radio" name="age" value="good" id="good" /> <label for="good">Excelent</label><br />
       <input type="radio" name="age" value="hight" id="hight" /> <label for="hight">Bien</label><br />
       <input type="radio" name="age" value="moyen" id="moyen" /> <label for="moyen">Moyen</label><br />
       <input type="radio" name="age" value="low" id="low" /> <label for="low">Boff !</label>
   </p>
   
   
    <p>Proposer une reponse : <input type="text" name="reponse" placeholder="Ex : proposer une reponse" /></p>

 <p><input type="submit" value="OK"></p>
</form>

<?php
include("JSON.php");

//**********************************************************************************************
//   Vérifit si une clée existe dans un tableau :
function existsoarray(){
	$search_array = array('premier' => 1, 'second' => 4);
	if (array_key_exists('premier', $search_array)) {
		echo "L'élément 'premier' existe dans le tableau";
	}
}
//-------------------
function array_random($arr, $num = 1) {
    shuffle($arr);
   
    $r = array();
    for ($i = 0; $i < $num; $i++) {
        $r[] = $arr[$i];
    }
    return $num == 1 ? $r[0] : $r;
}
//Exemple #1 Exemple avec array_rand()
function randoarray(){
	$input = array("Neo", "Morpheus", "Trinity", "Cypher", "Tank");
	$rand_keys = array_rand($input, 2);
	echo $input[$rand_keys[0]] . "\n";
	echo $input[$rand_keys[1]] . "\n";
}
//Exemple #1 Exemple avec array_search()
function searchoarray(){
	$array = array(0 => 'blue', 1 => 'red', 2 => 'green', 3 => 'red');
	$key = array_search('green', $array); // $key = 2;
	$key = array_search('red', $array);   // $key = 1;
}
//Exemple #1 Exemple avec array_push()
//array_push — Empile un ou plusieurs éléments à la fin d'un tableau
function addtoarray(){
	$stack = array("orange", "banana");
	array_push($stack, "apple", "raspberry");
	print_r($stack);
}

//**********************************************************************************************
//********************************* function outils ********************************************
//**********************************************************************************************

/**
Affichage json
*/
function affichagerow ($json){
  echo '<textarea name="ameliorer" id="ameliorer" rows="3" cols="50">';
  //print_r($json);
  $arrayOfObjsnew = $json->menuitem;
  foreach ($arrayOfObjsnew as $key => $object) {
      echo $object->entrer.'-';
      echo $object->result.'-';
      echo $object->poid;
      echo ' ';
  }
  echo '</textarea>';
}

/**
recherche in object
*/
function rechercherow ($json_a, $value){
  //affichagerow ($json_a);
  $input = array("-_-");
  $arrayOfObjsnew = $json_a->menuitem;
  foreach ($arrayOfObjsnew as $key => $object) {
    //strtolower ($newsujet);
      if (strtolower ($object->entrer) == strtolower ($value)){
	    array_push($input, $object->result);
	  }
      //echo $object->result.'-';
      //echo $object->poid;	  
  }
  print_r($input);
  // résultat aléatoire:
  	//$rand_keys = array_rand($input);
	//print_r(array_random($input));
	//echo array_random($input) . "\n";
  return array_random($input);
}

/**
retourne un tableau depuis le fichier jscon
*/
function loadvaluejson () {
$json = new Services_JSON();
echo 'Load Json<br>';
//$string = file_get_contents("results.json");
//$json_a=json_decode($string,true);
$input = file_get_contents('results2.json', 1000000);
$json_a = $json->decode($input);

affichagerow ($json_a);

echo '<br>';
return $json_a;
}



/**
Enregistre tableau dans le fichier jscon
*/
function savevaluejson ($arr) {
	$json = new Services_JSON();
	$fp = fopen('results2.json', 'w');
	$output = $json->encode($arr);
	fwrite($fp, $output);
	fclose($fp);
}

//***********************************************************************************************
//************************************ function principal ***************************************
//***********************************************************************************************


function reseau ($json,$entrer,$poid){
$resultat = 'Bonjour';
if ( $entrer != null){
	if ( $poid != null){
		//echo ' on propose un resultat en fonction du poid';
		//resultat = List [poid];
		$resultat = 'Bonjour';
		if($resultat == null){
			echo ' on propose une valeur aleatoire.';
			$resultat = 'Valeur aleatoire.';
		}
	}else{
		//echo ' on donne le resultat qui a le plus de poid.';
	   //on donne le resultat qui a le plus de poid.
	   $resultat = rechercherow ($json, $entrer);
	   
	}
}
return $resultat;
}

function appel ($entrer, $resul, $good){
         $json_a = loadvaluejson ();
         $poid = 0;
		 
		 //Enregistrement si correctif
		 
		     if ( $good == 'good') {
      $poid = 100;
      enregistrer ($json_a , $entrer, $resul, $poid);
    } else if ( $good == 'hight'){
      $poid = 75;
      enregistrer ($json_a , $entrer, $resul, $poid);
    } else if ( $good == 'moyen'){
      $poid = 50;
      enregistrer ($json_a , $entrer, $resul, $poid);
    } else if ( $good == 'low'){
      $poid = 25;
      enregistrer ($json_a , $entrer, $resul, $poid);
    }else{
      $poid = 0;
      //enregistrer ($json_a , $entrer, $resul, $poid);
   }
		 
        // if ($resul == null){
            $resul = reseau($json_a ,$entrer, $poid);
        // }


return $resul;
}

function enregistrer ($json_a, $entrer, $resul, $poid){
    //Listentrer.add (entrer, resul, poid);
    echo 'Enregistement de ('.$entrer.','. $resul.','. $poid.');<br>';
   // On ajout un object
	$obj = new stdClass;
	$obj2 = new stdClass;
	$obj2->entrer = $entrer;
	$obj2->result = $resul;
	$obj2->poid = $poid;
	$obj->menuitem = array($obj2);

	//$extended = (object) array_merge((array)$json_a, (array)$obj);
	$extended =  (object) array_merge_recursive ( (array) $json_a, (array) $obj );

	affichagerow ($extended);
    savevaluejson ($extended);
}

function char($text)
{
	$text = htmlentities($text, ENT_NOQUOTES, "UTF-8");
	$text = htmlspecialchars_decode($text);
	return $text;
}

function charspe($text)
{
    // Suppression des caractères spéciaux:
  $interdit=array("“","¯","Â","&nbsp;","_",">", "<","«","»", ":", "*","\\", "/", "|", "}", "{", ")", "[", "]", "(","$","#","~","ª","´","¨","¢","©",".","=","%", "\"",";","…","²");
  $reponse = str_replace($interdit, " ", $text);
  return $text;
}



//programme principal ***************************************************************
//init
if (isset($_POST['nom'])){
   $entrer = htmlspecialchars($_POST['nom']);
   $entrer = charspe($entrer);
}else{
   $entrer = "Bonjour";
}
// CheckBox :
if (isset($_POST['age'])){
   $good = $_POST['age'];
   //echo 'valeur >>     '.$good;

}else{
   $good = "Bonjour";
}
// Correction
if (isset($_POST['reponse'])){
   $newsujet = $_POST['reponse'];
   $reponse = charspe($newsujet);

}else{
   $reponse = "Bonjour";
}

// Appel principale

if ($good != null){
$result = appel($entrer,$reponse, $good);

}

echo '<br><br><h1>'.charspe($result).'</h1>';

?>