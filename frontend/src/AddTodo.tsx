import { Button, HStack, Input, useToast } from '@chakra-ui/react';
import React, { useState } from 'react';
import { nanoid } from 'nanoid';
import { Todo } from './services/BackendService';

interface AddTodoProps {
    addTodo: any
}

export const AddTodo = (props: AddTodoProps) => {

    const toast = useToast();

    const handleSubmit = (e: any) => {
        e.preventDefault();
        if (!content) {
            toast({
                title: 'No content',
                status: 'error',
                duration: 2000,
                isClosable: true,
            });
            return;
        }

        const todo:Todo = {
            id: nanoid(),
            body: content,
        };

        props.addTodo(todo);
        setContent('');
    }

    const [content, setContent] = useState('');

    return (
        <form onSubmit={handleSubmit}>
            <HStack mt='8'>
                <Input
                    variant='filled'
                    placeholder='add some todos'
                    value={content}
                    aria-label="todo-input" 
                    onChange={(e) => setContent(e.target.value)}
                />
                <Button className="submit-btn" colorScheme='blue' px='8' type='submit'>Add</Button>
            </HStack>
        </form>
    );
}

export default AddTodo;
