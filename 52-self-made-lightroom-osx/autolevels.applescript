# This script automates process of color adjustment using Preview.app, make it batch
# Based on http://stackoverflow.com/questions/31750123/using-applescript-to-adjust-colors-in-preview

set imagesFolder to (choose folder with prompt "Select the start folder")

try
	my process_folder(imagesFolder)
on error errStr number errorNumber
	display dialog errStr
end try

on process_folder(imagesFolder)
	tell application "Finder" to set allImages to every item of folder imagesFolder

	set adjustColorWindow to missing value
	repeat with anImage in allImages
		repeat 1 times
			if name of anImage is ".DS_Store" then exit repeat

			tell application "Finder" to open anImage
			activate application "Preview"
			tell application "System Events"
				tell process "Preview"
					repeat until exists (1st window whose value of attribute "AXSubRole" is "AXStandardWindow")
						delay 0.2
					end repeat

					set documentWindow to (name of 1st window whose value of attribute "AXSubRole" is "AXStandardWindow")
					if adjustColorWindow is missing value then
						click menu item "Adjust ColorÉ" of menu 1 of menu bar item "Tools" of menu bar 1
						repeat until exists (1st window whose title starts with "Adjust Color")
							delay 0.2
						end repeat
					end if

					set adjustColorWindow to (1st window whose title starts with "Adjust Color")
					tell adjustColorWindow
						click button "Auto Levels"
					end tell

					click menu item "Save" of menu 1 of menu bar item "File" of menu bar 1
					click button 1 of window documentWindow
				end tell
			end tell
		end repeat
	end repeat
end process_folder
