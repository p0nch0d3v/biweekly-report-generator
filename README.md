 # Biweekly Report Generator

The intention is facilitate the generation of the report based on a daily log, 
as example:
```md
- 2024/01/01
    - project A
        - task A
    - project C
        - task D
            - this third tab would be ignored

- 2024/01/02
    - project A
        - task A
    - project B
        - task E
            - this third tab would be ignored
```

Where:
> Project is identified with a single tab (or four spaces)

> Task is identified with a double tab (or eight spaces)

 Usage in windows, open a terminal such powershell and execute:
 ```bash
./biweekly-report-generator-linux-amd64 file.md
 ```
or in Linux/MacOs or any *nix OS, open a terminal and execute:
```powershell
.\biweekly-report-generator-windows-amd64.exe file.txt
```

Output:
```
- project A
- project B
- project C

- project A - task A
- project B - task E
- project C - task D
```
