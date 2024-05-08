<h1><strong><span style="font-size:23pt;">Vowel Mirroring Program</span></strong></h1>
<p><span style="font-size:11pt;">This program checks the arguments for vowels and mirrors the position of the vowels in the arguments. If the arguments contain vowels (excluding &apos;y&apos;), the program mirrors the position of the vowels in the arguments. If the arguments do not have any vowels, the program just prints the arguments. If the number of arguments is less than 1, the program displays a new line.</span></p>
<h2><strong><span style="font-size:17pt;">Usage</span></strong></h2>

<p><span style="font-size:11pt;">The program can be executed with the following command:</span></p>

<pre>
<span style="font-size:11pt;">bash</span>
<span style="color:#188038;font-size:11pt;">go run . [arguments]</span>
</pre>
<p><b><span style="font-size:13pt;">Examples</span></b></p>
<hr>
<p><br></p>
<pre>
<span style="font-size:11pt;">bash</span>
<span style="color:#188038;font-size:11pt;">$ go run . &quot;Hello World&quot; | cat -e</span>
<p><span style="color:#188038;font-size:11pt;">Hollo Werld$</span></p>
<span style="color:#188038;font-size:11pt;">$ go run . &quot;HEllO World&quot; &quot;problem solved&quot;</span>
<p><span style="color:#188038;font-size:11pt;">Hello Werld problom sOlvEd</span></p>
<span style="color:#188038;font-size:11pt;">$ go run . &quot;str&quot; &quot;shh&quot; &quot;psst&quot;</span>
<p><span style="color:#188038;font-size:11pt;">str shh psst</span></p>
<span style="color:#188038;font-size:11pt;">$ go run . &quot;happy thoughts&quot; &quot;good luck&quot;</span>
<p><span style="color:#188038;font-size:11pt;">huppy thooghts guod lack</span></p>
<span style="color:#188038;font-size:11pt;">$ go run . &quot;aEi&quot; &quot;Ou&quot;</span>
<p><span style="color:#188038;font-size:11pt;">uOi Ea</span></p>
<span style="color:#188038;font-size:11pt;">$ go run .</span>
<span style="color:#188038;font-size:11pt;">$</span>
</pre>
<h2><strong><span style="font-size:17pt;">Implementation Details</span></strong></h2>

<ul>
    <li style="list-style-type:disc;font-size:11pt;">
        <p><span style="font-size:11pt;">The program processes command-line arguments to check for vowels and mirror their positions if found.</span></p>
    </li>
    <li style="list-style-type:disc;font-size:11pt;">
        <p><span style="font-size:11pt;">It maintains a list of vowels (excluding &apos;y&apos;) and scans the input string for vowel occurrences.</span></p>
    </li>
    <li style="list-style-type:disc;font-size:11pt;">
        <p><span style="font-size:11pt;">When a vowel is found, it mirrors its position in the string and constructs the mirrored string.</span></p>
    </li>
    <li style="list-style-type:disc;font-size:11pt;">
        <p><span style="font-size:11pt;">If no arguments are provided, the program prints a new line (&quot;\n&quot;).</span></p>
    </li>
</ul>
<h2><strong><span style="font-size:17pt;">Customization</span></strong></h2>

<p><span style="font-size:11pt;">Feel free to customize the program according to your needs. You can modify the code to handle additional cases or improve the efficiency of vowel checking and mirroring.</span></p>

<h2><strong><span style="font-size:17pt;">License</span></strong></h2>

<p><span style="font-size:11pt;">This program is open-source and available under the MIT License. Feel free to use, modify, and distribute it for any purpose. Contributions and feedback are welcome!</span></p>

<h2><strong><span style="font-size:17pt;">Author</span></strong></h2>

<p><span style="font-size:11pt;">By David Jesse Odhiambo</span></p>
<p><span style="font-size:11pt;">Apprentice SoftwareDeveloper, Zone01 Kisumu</span></p>
<p><br></p>
<hr>
