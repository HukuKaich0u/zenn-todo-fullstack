"use client";

import { ModeToggle } from "@/components/mode-toggle";
import { useState, useEffect } from "react";

type Todo = {
  id: number;
  title: string;
  iscompleted: boolean;
};

const Home = () => {
  const [todos, setTodos] = useState<Todo[]>([]);
  const [newTitle, setNewTitle] = useState("");

  useEffect(() => {
    fetchTodos();
  }, []);

  const fetchTodos = async () => {
    const res = await fetch("http://localhost:8080/todos");
    const data = await res.json();
    setTodos(data);
  };

  const handleToggleCompleted = async (todo: Todo) => {
    await fetch(`http://localhost:8080/todos/${todo.id}`, {
      method: "PUT",
      headers: { "Content-Type": "application/json" },
      body: JSON.stringify({
        title: todo.title,
        iscompleted: !todo.iscompleted,
      }),
    });
    fetchTodos();
  };

  const handleAddTodo = async () => {
    await fetch(`http://localhost:8080/todos`, {
      method: "POST",
      headers: { "Content-Type": "application/json" },
      body: JSON.stringify({
        title: newTitle,
        iscompleted: false,
      }),
    });
    setNewTitle("");
    fetchTodos();
  };
  return (
    <>
      <div className="m-5">
        <ModeToggle />
      </div>
      <div className="p-6 max-w-2xl mx-auto">
        <h1 className="text-2xl font-semibold text-center mb-6">
          📋 Todo List
        </h1>

        <div className="flex items-center gap-2 mb-6">
          <input
            type="text"
            value={newTitle}
            onChange={(e) => setNewTitle(e.target.value)}
            placeholder="新しいTodoを入力してください"
            className="flex p-2 rounded bg-gray-700 placeholder-gray-400"
          />
          <button
            onClick={handleAddTodo}
            className="bg-blue-500 text-white px-4 py-2 rounded hover:bg-blue-600 transition"
          >
            追加
          </button>
        </div>

        {/* Todo一覧 */}
        <ul className="space-y-4">
          {todos.map((todo) => (
            <li
              key={todo.id}
              className="flex justify-between items-center p-4 bg-gray-800 border border-gray-700 rounded-xl shadow-sm hover:shadow-md transition"
            >
              <div className="flex items-center">
                <input
                  type="checkbox"
                  checked={todo.iscompleted}
                  onChange={() => handleToggleCompleted(todo)}
                  className="w-5 h-5 mr-4 accent-green-600"
                />
                <span className="text-lg font-medium text-white">
                  {todo.title}
                </span>
              </div>
            </li>
          ))}
        </ul>
      </div>
    </>
  );
};

export default Home;
