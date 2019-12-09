# architecture-lab-3

Максимальна кількість потоків, які можуть виконуватись паралельно, дорівнює кількості логічних процесорів на машині. Це значення дорівнює **кількість сокетів * кількість ядер в кожному сокеті * кількість гіперпотоків кожного ядра**.

У нашому випадку це значення дорівнює 4 на машині з Windows:
![alt text](https://github.com/dkalytovskyi/architecture-lab-3/blob/master/markdown-images/windows.png "Windows stats")

Аналогічно на машині з Linux:
![alt text](https://github.com/dkalytovskyi/architecture-lab-3/blob/master/markdown-images/linux.png "Linux stats")