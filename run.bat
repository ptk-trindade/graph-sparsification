@echo off
setlocal EnableDelayedExpansion

rem Run the alghorithm for the graphs
set iValues= 0002 001 002 003 004 005 006 007 008 009 01 02 05

rem echo avgDegree,stdDev,minDegree,maxDegree,components,smallest comp,biggest comp,EBS EBD rs,EBS EBD p-value,EBS EBD jcc,EBD ER rs,EBD ER p-value,EBD ER jcc,EBS ER rs,EBS ER p-value,EBS ER jcc > output_bat.txt

for %%i in (%iValues%) do (
    rem Loop through each j value

    echo Running for i = %%i >> output_bat.txt
    
    for /L %%j in (0,1,9) do (
        rem Construct the file path
        set filePath=txtFiles/inputs/random/1000_%%i/1000_%%i_%%j.txt

        rem Check if the file exists
        if exist !filePath! (
            rem Run the command and append the output to output_bat.txt
            powershell -Command "Get-Content !filePath! | .\main.exe >> output_bat.txt"
        ) else (
            echo File !filePath! does not exist.
        )
    )
)



@REM rem create random graphs
@REM echo 1000 0.002 > temp_input.txt

@REM rem recreate folder
@REM set folderPath=txtFiles/inputs/random/1000_0002

@REM rmdir /q /s "!folderPath!"
@REM mkdir "!folderPath!"

@REM for /L %%j in (0,1,9) do (

@REM     rem Construct the file path
@REM     set filePath=!folderPath!/1000_0002_%%j.txt

@REM     powershell -Command "Get-Content temp_input.txt | .\graphGenerator.exe > !filePath!"
@REM )

@REM echo 1000 0.01 > temp_input.txt

@REM rem recreate folder
@REM set folderPath=txtFiles/inputs/random/1000_001

@REM rmdir /q /s "!folderPath!"
@REM mkdir "!folderPath!"

@REM for /L %%j in (0,1,9) do (

@REM     rem Construct the file path
@REM     set filePath=!folderPath!/1000_001_%%j.txt

@REM     powershell -Command "Get-Content temp_input.txt | .\graphGenerator.exe > !filePath!"
@REM )

@REM echo 1000 0.02 > temp_input.txt

@REM rem recreate folder
@REM set folderPath=txtFiles/inputs/random/1000_002

@REM rmdir /q /s "!folderPath!"
@REM mkdir "!folderPath!"

@REM for /L %%j in (0,1,9) do (

@REM     rem Construct the file path
@REM     set filePath=!folderPath!/1000_002_%%j.txt

@REM     powershell -Command "Get-Content temp_input.txt | .\graphGenerator.exe > !filePath!"
@REM )

@REM echo 1000 0.03 > temp_input.txt

@REM rem recreate folder
@REM set folderPath=txtFiles/inputs/random/1000_003

@REM rmdir /q /s "!folderPath!"
@REM mkdir "!folderPath!"

@REM for /L %%j in (0,1,9) do (

@REM     rem Construct the file path
@REM     set filePath=!folderPath!/1000_003_%%j.txt

@REM     powershell -Command "Get-Content temp_input.txt | .\graphGenerator.exe > !filePath!"
@REM )

@REM echo 1000 0.04 > temp_input.txt

@REM rem recreate folder
@REM set folderPath=txtFiles/inputs/random/1000_004

@REM rmdir /q /s "!folderPath!"
@REM mkdir "!folderPath!"

@REM for /L %%j in (0,1,9) do (

@REM     rem Construct the file path
@REM     set filePath=!folderPath!/1000_004_%%j.txt

@REM     powershell -Command "Get-Content temp_input.txt | .\graphGenerator.exe > !filePath!"
@REM )

@REM echo 1000 0.05 > temp_input.txt

@REM rem recreate folder
@REM set folderPath=txtFiles/inputs/random/1000_005

@REM rmdir /q /s "!folderPath!"
@REM mkdir "!folderPath!"

@REM for /L %%j in (0,1,9) do (

@REM     rem Construct the file path
@REM     set filePath=!folderPath!/1000_005_%%j.txt

@REM     powershell -Command "Get-Content temp_input.txt | .\graphGenerator.exe > !filePath!"
@REM )

@REM echo 1000 0.06 > temp_input.txt

@REM rem recreate folder
@REM set folderPath=txtFiles/inputs/random/1000_006

@REM rmdir /q /s "!folderPath!"
@REM mkdir "!folderPath!"

@REM for /L %%j in (0,1,9) do (

@REM     rem Construct the file path
@REM     set filePath=!folderPath!/1000_006_%%j.txt

@REM     powershell -Command "Get-Content temp_input.txt | .\graphGenerator.exe > !filePath!"
@REM )

@REM echo 1000 0.07 > temp_input.txt

@REM rem recreate folder
@REM set folderPath=txtFiles/inputs/random/1000_007

@REM rmdir /q /s "!folderPath!"
@REM mkdir "!folderPath!"

@REM for /L %%j in (0,1,9) do (

@REM     rem Construct the file path
@REM     set filePath=!folderPath!/1000_007_%%j.txt

@REM     powershell -Command "Get-Content temp_input.txt | .\graphGenerator.exe > !filePath!"
@REM )

@REM echo 1000 0.08 > temp_input.txt

@REM rem recreate folder
@REM set folderPath=txtFiles/inputs/random/1000_008

@REM rmdir /q /s "!folderPath!"
@REM mkdir "!folderPath!"

@REM for /L %%j in (0,1,9) do (

@REM     rem Construct the file path
@REM     set filePath=!folderPath!/1000_008_%%j.txt

@REM     powershell -Command "Get-Content temp_input.txt | .\graphGenerator.exe > !filePath!"
@REM )

@REM echo 1000 0.09 > temp_input.txt

@REM rem recreate folder
@REM set folderPath=txtFiles/inputs/random/1000_009

@REM rmdir /q /s "!folderPath!"
@REM mkdir "!folderPath!"

@REM for /L %%j in (0,1,9) do (

@REM     rem Construct the file path
@REM     set filePath=!folderPath!/1000_009_%%j.txt

@REM     powershell -Command "Get-Content temp_input.txt | .\graphGenerator.exe > !filePath!"
@REM )

@REM echo 1000 0.1 > temp_input.txt

@REM rem recreate folder
@REM set folderPath=txtFiles/inputs/random/1000_01

@REM rmdir /q /s "!folderPath!"
@REM mkdir "!folderPath!"

@REM for /L %%j in (0,1,9) do (

@REM     rem Construct the file path
@REM     set filePath=!folderPath!/1000_01_%%j.txt

@REM     powershell -Command "Get-Content temp_input.txt | .\graphGenerator.exe > !filePath!"
@REM )

@REM echo 1000 0.2 > temp_input.txt

@REM rem recreate folder
@REM set folderPath=txtFiles/inputs/random/1000_02

@REM rmdir /q /s "!folderPath!"
@REM mkdir "!folderPath!"

@REM for /L %%j in (0,1,9) do (

@REM     rem Construct the file path
@REM     set filePath=!folderPath!/1000_02_%%j.txt

@REM     powershell -Command "Get-Content temp_input.txt | .\graphGenerator.exe > !filePath!"
@REM )

@REM echo 1000 0.5 > temp_input.txt

@REM rem recreate folder
@REM set folderPath=txtFiles/inputs/random/1000_05

@REM rmdir /q /s "!folderPath!"
@REM mkdir "!folderPath!"

@REM for /L %%j in (0,1,9) do (

@REM     rem Construct the file path
@REM     set filePath=!folderPath!/1000_05_%%j.txt

@REM     powershell -Command "Get-Content temp_input.txt | .\graphGenerator.exe > !filePath!"
@REM )

endlocal