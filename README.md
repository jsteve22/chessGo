# chessGo

The goal of this repository is to create a chess game and a chess engine in go. The code to hold the chess game will be in the board package.

Contributors: Jeremy Stevens


#### Building

1. Clone the EmpDocker git repository by running:
    ```
    git clone https://github.com/jsteve22/chessGo.git
    ```

2. Enter the Framework directory: `cd chessGo/`

3. With docker installed, run:
    ```
    docker build -t chessgo .
    docker run -it --name chess --mount "type=bind,source=$PWD,target=/home/" -p 8080:80 chessgo
    ```
    and re-run it with 
    ```
    docker start -i chess
    ```