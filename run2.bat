powershell -Command "Get-Content .\txtFiles\inputs\ABCommunityD\abcd_graph_10.txt | go run .\cmd\main.go > results_abcd.txt"
powershell -Command "Get-Content .\txtFiles\inputs\erdosRenyi\erdosRenyi_4000.txt | go run .\cmd\main.go > results_erdos.txt"
powershell -Command "Get-Content .\txtFiles\inputs\real_graphs\facebook.txt | go run .\cmd\main.go > results_fb.txt"
powershell -Command "Get-Content .\txtFiles\inputs\real_graphs\CA-GrQc.txt | go run .\cmd\main.go > results_collab.txt"