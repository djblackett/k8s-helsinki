import React from "react";
import { useDispatch, useSelector } from "react-redux";
import { selectColorMode } from "./features/colorMode/colorModeSlice";
import {addListItem, addTodo} from "./features/listItems/listItemsSlice";

function InputBar() {
  const dispatch = useDispatch();
  let mode = useSelector(selectColorMode);

  const handleClick = async () => {
    const input = document.getElementById("input");
    console.log(input)
    let text = input.value
    if (text === "") return;
    const newEntry = {
      text: text,
      completed: false,
    };

    const res = await addTodo(newEntry)
    console.log(res)
    dispatch(addListItem(res));
    document.getElementById("input").value = "";
  }

  const handleEnterPress = async (event) => {
    if (event.key === "Enter") {
      let text = event.target.value;
      if (text === "") return;
      const newEntry = {
        text: text,
        completed: false,
      };

      const res = await addTodo(newEntry)
      console.log(res)
      dispatch(addListItem(res));
      document.getElementById("input").value = "";
    }
  };

  return (
    <div
      id="input-component"
      className={`input-component-${mode}`}
      tabIndex={-1}
    >
      <div id="outer-circle">
        <div id="circle" className={`circle-${mode}`}></div>
      </div>
      <input
        id="input"
        className={`input-${mode}`}
        type="text"
        placeholder="Create a new todo..."
        onKeyDown={(e) => handleEnterPress(e)}
      />
      <button id="send-button" className={`send-${mode}`} onClick={() => handleClick()}>Send</button>
    </div>
  );
}

export default InputBar;
