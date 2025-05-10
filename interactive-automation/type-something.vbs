' some code to simulate typing in windows using visual basic script
Option Explicit

Dim sh

Sub out( msg )
	wscript.stdout.write msg
End Sub

Sub echo( msg )
	out msg
	wscript.stdout.writeBlankLines 1
End Sub

Sub wait( ms )
	wscript.sleep ms
End Sub

Sub kcmd( s )
	sh.sendkeys s
End Sub

Sub ktype( s, min_time, max_time )
	Dim i, l, t, range_plus_one, ch
	l = len(s)
	range_plus_one = Int( max_time - min_time + 1 )
	for i = 1 to l
		ch = mid(s, i, 1)
		Select Case ch
		case "("
			ch = "{(}"
		End Select
		sh.sendkeys ch
		t = Int( min_time + Rnd() * range_plus_one )
		wait t
	next
End Sub

Sub main
	Set sh = CreateObject("wscript.shell")

	wait 3500
	echo "- SCRIPT START -"
	'kcmd "%{TAB}"
	'wait 2000
	'kcmd "{ESC}"
	'wait 500
	ktype "The words you want to type", 100, 200
	wait 200
	kcmd "{ENTER}"

	echo "- SCRIPT END -"
End Sub

main

