 # Biweekly Report Generator

The intention is facilitate the generation of the report based on a daily log, 
as example:
```md
- 2024/01/01
    - project A
        - task A
    - project C
        - task D

- 2024/01/02
    - project A
        - task A
    - project B
        - task E
```
Where:
> Project is identified with a single tab

> Task is identified with a double tab

 Usage:
 ```bash
./biweekly-report-generator-linux-amd64 file.md
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