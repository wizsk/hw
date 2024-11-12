$ErrorActionPreference = "Stop"  # Stop on errors


$progName = "HansWehrsDictionary"
$programDataPath = [System.Environment]::GetFolderPath('ProgramFiles')
$hwPath = Join-Path -Path $programDataPath -ChildPath $progName
$url = "https://github.com/wizsk/hw/releases/latest/download"
$sysTmp = [System.IO.Path]::GetTempPath()
$file = "hw_windows_x86_64.zip"
$downFilePath = Join-Path -Path $sysTmp -ChildPath $file

try {
    Write-Host "Downloading: $url/$file"
    Invoke-WebRequest -Uri "$url/$file" -OutFile $downFilePath
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

Write-Host ""
Write-Host "Update compleated"
Write-Host ""


# Removing files
try {
    Remove-Item -Path "$exFolder" -Recurse
    Remove-Item -Path $downFilePath

}
catch {}
