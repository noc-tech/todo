import { cleanup, fireEvent } from "@testing-library/react";
import React from "react";
import TodoList from "./TodoList";
import { render } from "./test-utils";

afterEach(cleanup);

const defaultProps = {
  todos: [{id:'id1', body:'get some milk'}],
  deleteTodo: jest.fn(),
};

test('display no todos if there is nothing todo', () => {
  const { queryByText } = render(<TodoList {...defaultProps} todos={[]} />);
  expect(queryByText("No todos")).toBeTruthy();
});

test('todo list render correctly', () => {
  const { queryByText } = render(<TodoList {...defaultProps} />);
  expect(queryByText("get some milk")).toBeTruthy();
});

test('calls correct function on remove', () => {
  const deleteTodo = jest.fn();
  const { getByText } = render(<TodoList {...defaultProps} deleteTodo={deleteTodo} />)
  fireEvent.click(getByText("remove"));
  expect(deleteTodo).toHaveBeenCalled();
});
