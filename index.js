const taskInput = document.getElementById('taskInput');
const addTaskButton = document.getElementById('addTask');
const taskList1 = document.getElementById('taskList1');
const taskList2 = document.getElementById('taskList2');
const taskList3 = document.getElementById('taskList3');

const url = 'http://localhost:8080/tasks';

function updateTaskStatusByListItem(listItem, newStatus) {
    if (listItem.dataset.taskStatus !== newStatus) {
        listItem.dataset.taskStatus = newStatus
        var task = {
            id: listItem.dataset.taskId,
            title: listItem.dataset.taskTitle,
            status: +listItem.dataset.taskStatus, // Из строки делаем число
        }
        // Обновляем таску
        updateTask(task)
            .catch(error => {
                // Обработка ошибок
                console.error('error updating task:', error.message);
            });
    }
}

function createTaskElement(task) {
    const listItem = document.createElement('li');
    listItem.innerHTML = `
        <span>${task.title}</span>
        <div>
        <button class="toTheBegining">В начало</button>
        <button class="inProcess">В процессе</button>
        <button class="ready">Готово</button>
        <button class="deleteTask">Удалить</button>
        </div>
    `;
    listItem.dataset.taskId = task.id
    listItem.dataset.taskTitle = task.title
    listItem.dataset.taskStatus = task.status

    const processButton = listItem.querySelector('.inProcess');
    processButton.addEventListener('click', () => {
        taskList2.appendChild(listItem);

        updateTaskStatusByListItem(listItem, '2')
    });

    const readyButton = listItem.querySelector('.ready');
    readyButton.addEventListener('click', () => {
        taskList3.appendChild(listItem);

        updateTaskStatusByListItem(listItem, '3')
    });

    const beginButton = listItem.querySelector('.toTheBegining');
    beginButton.addEventListener('click', () => {
        taskList1.appendChild(listItem);

        updateTaskStatusByListItem(listItem, '1')
    });

    const deleteButton = listItem.querySelector('.deleteTask');
    deleteButton.addEventListener('click', (event) => {
        const currentList = event.target.parentNode.parentNode.parentNode;
        if (currentList == taskList1) taskList1.removeChild(listItem);
        if (currentList == taskList2) taskList2.removeChild(listItem);
        if (currentList == taskList3) taskList3.removeChild(listItem);

        updateTaskStatusByListItem(listItem, '4')
    });

    switch (task.status) {
        case 1:
            taskList1.appendChild(listItem);
            break;
        case 2:
            taskList2.appendChild(listItem);
            break;
        case 3:
            taskList3.appendChild(listItem);
            break;
        default:
            console.log('invalid task status');
    }
}

function fetchTasks() {
    // Получаем токен из localStorage
    const token = localStorage.getItem('token');
    // Опции запроса
    var requestOptions = {
        method: 'GET',
        headers: {
            'Authorization': `Bearer ${token}`,
        },
    };
    // Отправляем GET-запрос на сервер
    return fetch(url, requestOptions)
        .then(response => {
            // Проверка на успешный ответ (код 200-299)
            if (!response.ok) {
                if (response.status === 401) {
                    // Нужно заново выполнить login. Поэтому переходим на страницу для login
                    window.location.href = 'auth.html';
                } else {
                    throw new Error(`Network response was not ok, status: ${response.status}`);
                }
            }
            // Преобразование ответа в JSON
            return response.json();
        })
        .catch(error => {
            // Обработка ошибок
            console.error('error fetching tasks:', error.message);
        });
}

function createTask(title, status) {
    // Новая таска
    var task = {
        title: title,
        status: status,
    };
    // Получаем токен из localStorage
    const token = localStorage.getItem('token');
    // Опции запроса
    var requestOptions = {
        method: 'POST',
        headers: {
            'Content-Type': 'application/json',
            'Authorization': `Bearer ${token}`,
        },
        body: JSON.stringify(task)
    };
    // Отправляем POST-запрос на сервер
    return fetch(url, requestOptions)
        .then(response => {
            // Проверка на успешный ответ (код 200-299)
            if (!response.ok) {
                if (response.status === 401) {
                    // Нужно заново выполнить login. Поэтому переходим на страницу для login
                    window.location.href = 'auth.html';
                } else {
                    throw new Error(`Network response was not ok, status: ${response.status}`);
                }
            }
            // Преобразование ответа в JSON
            return response.json();
        })
        .catch(error => {
            // Обработка ошибок
            console.error('There was a problem with the fetch operation:', error.message);
        });
}

function updateTask(task) {
    // Получаем токен из localStorage
    const token = localStorage.getItem('token');
    // Опции запроса
    var requestOptions = {
        method: 'PUT',
        headers: {
            'Content-Type': 'application/json',
            'Authorization': `Bearer ${token}`,
        },
        body: JSON.stringify(task)
    };
    // Отправляем POST-запрос на сервер
    return fetch(url, requestOptions)
        .then(response => {
            // Проверка на успешный ответ (код 200-299)
            if (!response.ok) {
                if (response.status === 401) {
                    // Нужно заново выполнить login. Поэтому переходим на страницу для login
                    window.location.href = 'auth.html';
                } else {
                    throw new Error(`Network response was not ok, status: ${response.status}`);
                }
            }
            // Преобразование ответа в JSON
            return response.json();
        })
        .catch(error => {
            // Обработка ошибок
            console.error('There was a problem with the fetch operation:', error.message);
        });
}

// Получаем все таски и добавляем на доску тасок
fetchTasks()
    .then(data => {
        console.log(data)
        // Перебираем элементы массива
        for (let task of data) {
            // Добавляем таску на доску тасок
            createTaskElement(task)
        }
    })
    .catch(error => {
        // Обработка ошибок
        console.error('There was an error outside fetchData:', error.message);
    });

addTaskButton.addEventListener('click', () => {
    const taskTitle = taskInput.value.trim();

    if (taskTitle === '') {
        alert('Пожалуйста, введите задачу.');
        return;
    }

    createTask(taskTitle, 1)
        .then(task => {
            // Добавляем таску на общую доску тасок
            createTaskElement(task);
            // Очищаем поле ввода
            taskInput.value = '';
        })
        .catch(error => {
            // Обработка ошибок
            console.error('error creating task:', error.message);
        });
});
