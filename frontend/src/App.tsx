import * as React from "react"
import {
  ChakraProvider,
  Box,
  VStack,
  Grid,
  theme,
  Heading,
} from "@chakra-ui/react"
import { ColorModeSwitcher } from "./ColorModeSwitcher"
import AddTodo from "./AddTodo";
import TodoList from "./TodoList";
import BackendService, { Todo } from "./services/BackendService";

export const App = () => {

  const [todos, setTodos] = React.useState<Todo[]>([]);

  React.useEffect(() => {
    const fetchData = async () => {
      const result = await BackendService.getTodos();
      setTodos(result.data);
    };
    fetchData();
  }, []);

  function deleteTodo(id: string) {
    BackendService.deleteTodo(id);
    const newTodos = todos.filter((todo: Todo) => {
      return todo.id !== id;
    });
    setTodos(newTodos);
  }

  function addTodo(todo: Todo) {
    BackendService.createTodo(todo);
    setTodos([...todos, todo]);
  }

  return (
    <ChakraProvider theme={theme}>
      <Box textAlign="center" fontSize="xl">
        <Grid minH="100vh" p={3}>
          <ColorModeSwitcher justifySelf="flex-end" />
          <VStack spacing={8}>
            <VStack p={4}>
              <Heading mb='8' fontWeight='extrabold' size='xl'>Todo Application</Heading>
              <TodoList todos={todos} deleteTodo={deleteTodo} />
              <AddTodo addTodo={addTodo} />
            </VStack>
          </VStack>
        </Grid>
      </Box>
    </ChakraProvider>
  )
}
