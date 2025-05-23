from collections import deque

def read_maze(filename):
    #Читает лабиринт из файла и возвращает его как двумерный список
    with open(filename, 'r') as f:
        maze = [list(line.strip()) for line in f.readlines()]
    return maze

def find_start_end(maze):
    #Находит начальную и конечную точки в лабиринте
    start, end = None, None
    for i in range(len(maze)):
        for j in range(len(maze[i])):
            if maze[i][j] == 'S':
                start = (i, j)
            elif maze[i][j] == 'E':
                end = (i, j)
    return start, end

def bfs(maze, start, end):
    #Алгоритм поиска в ширину для нахождения пути
    queue = deque()
    queue.append(start)
    visited = {start: None} #Словарь для отслеживания посещенных клеток и их предков
    
    while queue:
        current = queue.popleft()
        if current == end:
            break  #Нашли выход
        
        #Проверяем все соседние клетки (вверх, вниз, влево, вправо)
        for di, dj in [(-1, 0), (1, 0), (0, -1), (0, 1)]:
            neighbor = (current[0] + di, current[1] + dj)
            
            #Проверяем, что соседняя клетка в пределах лабиринта
            if (0 <= neighbor[0] < len(maze)) and (0 <= neighbor[1] < len(maze[0])):
                #Проверяем, что клетка проходима и не посещалась
                if (maze[neighbor[0]][neighbor[1]] == ' ' or maze[neighbor[0]][neighbor[1]] == 'E') and neighbor not in visited:
                    visited[neighbor] = current  #Запоминает, откуда пришли
                    queue.append(neighbor)
    
    #Восстанавливаем путь от конца к началу
    path = []
    if end in visited:
        current = end
        while current != start:
            path.append(current)
            current = visited[current]
        path.reverse()  #Переворачиваем, чтобы путь был от начала к концу
    
    return path

def mark_path(maze, path):
    #Помечает путь в лабиринте
    for i, j in path:
        if maze[i][j] != 'E':  #Не помечает конечную точку
            maze[i][j] = '*'
    return maze

def write_maze(filename, maze, steps):
    #Записывает лабиринт с путем в файл
    with open(filename, 'w') as f:
        f.write(f"Steps: {steps}\n")
        for row in maze:
            f.write(''.join(row) + '\n')

def main():
    input_file = 'input.txt'
    output_file = 'output.txt'
    
    #Читаем лабиринт
    maze = read_maze(input_file)
    
    #Находим начальную и конечную точки
    start, end = find_start_end(maze)
    
    if not start or not end:
        print("Ошибка: лабиринт должен содержать начальную (S) и конечную (E) точки")
        return
    
    #Ищем путь
    path = bfs(maze, start, end)
    
    if not path:
        print("Путь от S до E не найден")
        return
    
    #Помечаем путь и сохраняем результат
    marked_maze = mark_path(maze, path)
    write_maze(output_file, marked_maze, len(path))
    print(f"Путь найден, сохранено в {output_file}")

main()
