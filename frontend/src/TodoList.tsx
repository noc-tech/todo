import React from 'react';
import {
    HStack,
    VStack,
    Text,
    StackDivider,
    Spacer,
    Badge,
    Button,
} from '@chakra-ui/react';
import { Todo } from './services/BackendService';

interface TodoListProps {
    todos: Todo[]
    deleteTodo: any
}

export const TodoList = (props: TodoListProps) => {
    if (!props.todos.length) {
        return (
            <Badge colorScheme='green' p='4' m='4' borderRadius='lg'>
                No todos
            </Badge>
        );
    }

    return (
        <VStack
            divider={<StackDivider />}
            borderColor='gray.100'
            borderWidth='1px'
            p='2'
            borderRadius='lg'
            w='100%'
            maxW={{ base: '90vw', sm: '80vw', lg: '50vw', xl: '40vw' }}
            alignItems='stretch'
        >
            {props.todos.map((todo: Todo) => (
                <HStack key={todo.id}>
                    <Text onClick={() => props.deleteTodo(todo.id)}>{todo.body}</Text>
                    <Spacer />
                    <Button size='xs' colorScheme='red' px='2' onClick={() => props.deleteTodo(todo.id)}>remove</Button>
                </HStack>
            ))}
        </VStack>
    );
}

export default TodoList;
