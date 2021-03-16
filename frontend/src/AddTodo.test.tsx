import { cleanup, fireEvent } from "@testing-library/react";
import React from "react";
import { act } from "react-dom/test-utils";
import AddTodo from "./AddTodo";
import { render } from "./test-utils";

afterEach(cleanup);

const defaultProps = {
  addTodo: jest.fn(),
};

test('add button renders with correct text', () => {
  const { queryByText } = render(<AddTodo {...defaultProps} />);
  expect(queryByText("Add")).toBeTruthy();
});

test('calls correct function on submit', () => {
  const addTodo = jest.fn();
  const { getByText,getByPlaceholderText } = render(<AddTodo {...defaultProps} addTodo={addTodo} />)
  const input = getByPlaceholderText("add some todos");
  fireEvent.change(input, { target: { value: "get some milk" }});
  fireEvent.click(getByText('Add'));
  expect(addTodo).toHaveBeenCalled();
});
