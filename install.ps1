$ErrorActionPreference = "Stop"  # Stop on errors


# # Remove-ItemProperty -Path "HKCU:\Software\Microsoft\Windows\CurrentVersion\Run" -Name $progName -ErrorAction SilentlyContinue
# Write-Host "$programName removed from startup for all users."
#
# return
# # Add registry key for all users (requires admin privileges)
# Set-ItemProperty -Path "HKLM:\Software\Microsoft\Windows\CurrentVersion\Run" -Name $progName -Value $hwPath
# Write-Host "$progName added to startup for all users."

# Check if the folder exists
if (-Not (Test-Path -Path $hwPath)) {
    New-Item -Path $hwPath -ItemType Directory
    Write-Host "Directory created: $hwPath"
} else {
    Write-Host "Directory already exists: $hwPath"
}

try {
    Write-Host "Downloading: $url/$file"
    # Invoke-WebRequest -Uri "$url/$file" -OutFile $downFilePath
    Write-Host "Downloading: $icoUrl"
    # Invoke-WebRequest -Uri "$icoUrl" -OutFile $icoPath
}
catch {
    Write-Host "Err: Could Downlaoing failed! bye"
    return
}


$exFolder = Join-Path -Path $sysTmp -ChildPath $progName

try {
    Write-Host "Extracting: $downFilePath"
    Expand-Archive -Path $downFilePath -DestinationPath $exFolder -Force
}
catch {
    Write-Host "err: Extracting failed $downFilePath bye"
}


try {
    Write-Host "Copying exe to $hwPath"
    $hwExePath = Join-Path -Path $exFolder -ChildPath "hw.exe"
    Copy-Item -Path $hwExePath -Destination $hwPath
}
catch {
    Write-Host "Err: Copying $hwExePath to $hwPath failed. bye!"
    return
}

Write-Host "Adding $hwPath to SYSTEM PATH"
$systemPath = [System.Environment]::GetEnvironmentVariable("PATH", [System.EnvironmentVariableTarget]::Machine)
# Check if the bin folder is already in the PATH
if ($systemPath -notlike "*$hwPath*") {
    $newSystemPath = "$systemPath;$hwPath"
    [System.Environment]::SetEnvironmentVariable("PATH", $newSystemPath, [System.EnvironmentVariableTarget]::Machine)
    $env:PATH = [System.Environment]::GetEnvironmentVariable("PATH", [System.EnvironmentVariableTarget]::Machine)
}


# Removing files
try {
    Remove-Item -Path "$exFolder" -Recurse
    # Remove-Item -Path $downFilePath

}
catch {}


$progName = "HansWehrsDictionary"
$programDataPath = [System.Environment]::GetFolderPath('ProgramFiles')
$hwPath = Join-Path -Path $programDataPath -ChildPath $progName

# Define the URL of the file and the path where you want to save it
$icoUrl = "https://raw.githubusercontent.com/wizsk/hw/refs/heads/main/assets/pub/hw.ico"
$url = "http://10.0.0.110:8001"
# $url = "https://github.com/wizsk/hw/releases/latest/download"
$sysTmp = [System.IO.Path]::GetTempPath()
$file = "hw_windows_x86_64.zip"
$downFilePath = Join-Path -Path $sysTmp -ChildPath $file
$icoPath = Join-Path -Path $hwPath -ChildPath "hw.ico"


# Arguments for the executable (if any)
$arguments = ""
# Get the path to the current user's Desktop
$desktopPath = [System.IO.Path]::Combine($env:USERPROFILE, "Desktop")
# Path to save the shortcut
$shortcutPath = [System.IO.Path]::Combine($desktopPath, "$progName.lnk")
# Create a WScript.Shell COM object to create the shortcut
$wshShell = New-Object -ComObject WScript.Shell
$shortcut = $wshShell.CreateShortcut($shortcutPath)
# Set the target executable and arguments
$hwExeFullPath = Join-Path -Path $hwPath -ChildPath "hw.exe"
$shortcut.TargetPath = $hwExeFullPath
$shortcut.Arguments = $arguments
$shortcut.IconLocation = $icoPath
# Save the shortcut
$shortcut.Save()

$shortcutPath = [System.IO.Path]::Combine($hwPath, "$progName.lnk")
$wshShell = New-Object -ComObject WScript.Shell
$shortcut = $wshShell.CreateShortcut($shortcutPath)
# Set the target executable and arguments
$hwExeFullPath = Join-Path -Path $hwPath -ChildPath "hw.exe"
$shortcut.TargetPath = $hwExeFullPath
$shortcut.Arguments = $arguments
$shortcut.IconLocation = $icoPath
# Save the shortcut
$shortcut.Save()
Write-Host "Shortcut created on Desktop."

Write-Host ""
Write-Host "Installation compleaded! Now run 'hw'"

