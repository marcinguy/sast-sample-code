# sast-sample-code

Samples from [Snyk Code Checker](https://snyk.io)  tested with [Betterscan.io](https://betterscan.io)


Betterscan engines used:

|Language|Engine|
|---|---|	
|Python|Bandit|
|Javascript|semgrep (nodejsscan),fluidattacks|
|PHP|progpilot|
|Java|PMD, semgrep|
|GO|Gosec, gostaticcheck|
|c#|insidersec, fluidattacks|
|Ruby|Brakeman|

All Languages were supported by OpenAI GPT.


# Summary

Betterscan.io is a Security Automation Software, so any check/software can be added to it (also better than the current ones)

None of the vulnerabilities were missed by Betterscan, it actually found more (Quality issues)

Betterscan does however so much more, like binary scannig using YARA, SBOM, dependencies vulnerabilities (like Dependabot) etc, adding your own checkers etc.

Please keep in mind actual missing vulnerabilities in real life could be much lower. Those samples were targeted for Snyk Code Checker. Feel free to try with "Goat" projects.


|Amount|Language|
|---|---|	
|0|for Python|
|0| for Javascript|
|0| for PHP|
|0| for Java|
|0| for GO|
|0| for C#|
|0| for Ruby|


Shows findings on the left with Snyk Code and on the right with Betterscan for the same code. 

|Python||
|---|---|	
|Snyk|Betterscan|
|Code Injection|Code Injection|	
|SQL Injection|SQL Injection|
|Open Redirect|Open Redirect|
||Bind on all interfaces|
||Several Quality issues|


|Javascript||
|---|---|
|Snyk|	Betterscan|
|SQL Injection	|SQL Injection|
|Hardcoded secret	|Hardcoded secret|
|Open Redirect	|Open Redirect
|Insufficied post Message Validation||
||Several Quality issues|

|PHP||
|---|---|
|Snyk|	Betterscan|
|Cross-site Scripting (XSS)|	Cross-site Scripting (XSS)|
|SQL Injection|	SQL Injection|
|Open Redirect	|Open Redirect|
|Use of Hardcoded Credentials	|Use of Hardcoded Credentials	|
||Several Quality issues|

|Java||
|---|---|
|Snyk|	Betterscan|
|Cross-site Scripting (XSS)	|Cross-site Scripting (XSS)|
|Improper Neutralization of CRLF Sequences in HTTP Headers	|Improper Neutralization of CRLF Sequences in HTTP Headers|
|Open Redirect	|Open Redirect|
|Origin Validation Error||	

|GO||
|---|---|
|Snyk|	Betterscan|
|Cross-site Scripting (XSS)|	Timeouts warning|
|Open Redirect|Open Redirect|
|Clear Text Logging||
||Errors not checked|

|C#|	|
|---|---|
|Snyk|	Betterscan|
|Open Redirect|Open Redirect|	
|SQL Injection	|SQL Injection|
|Cross-site Scripting (XSS)|Cross-site Scripting (XSS)|
||Several Quality issues|

|Ruby||
|---|---|	
|Snyk|	Betterscan|
|Open Redirect|Open Redirect|	
