import axios from 'axios';

class TodoService {
  async getTodos() {
    return await axios.get("https://api-todoapp.noctech.uk/todos");
  }

  async createTodo(todo: Todo) {
    return await axios.post("https://api-todoapp.noctech.uk/todos", todo);
  }

  async deleteTodo(id: any) {
    return await axios.delete("https://api-todoapp.noctech.uk/todos/" + id);
  }
}

export interface Todo {
  id?: string;
  body: string;
}

export default new TodoService();