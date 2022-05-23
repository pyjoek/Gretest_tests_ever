from datetime import *
import pyttsx3
engine = pyttsx3.init()
times = datetime.now().strftime("%H:%M")
if times < "12:00":
	engine.say("good morning joel")
	engine.runAndWait()
elif times < "6:00":
	engine.say("good afternoon joel")
	engine.runAndWait()
elif times > "6:00":
	engine.say("good evening sir")
	engine.runAndWait()
