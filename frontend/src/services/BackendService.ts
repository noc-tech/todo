import axios from 'axios';

class TodoService {
  async getTodos() {
    return await axios.get("http://localhost:3001/todos");
  }

  async createTodo(todo: Todo) {
    return await axios.post("http://localhost:3001/todos", todo);
  }

  async deleteTodo(id: any) {
    return await axios.delete("http://localhost:3001/todos/" + id);
  }
}

export interface Todo {
  id?: string;
  body: string;
}

export default new TodoService();