import React from 'react';
import Announcement from './Announcement';
import Event from './Event';
import Holiday from './Holiday';
import Scripture from './Scripture';
import Member from './Member';

const Collection = (props) => {
  let element = ""
  if (props.selection === "announcement")
  {
    //TODO: add logic to get only some text from the announcement
  element = <div className="list-group-item">
    <p className="">{props.header}{props.text}</p>
  </div>
  } else if (props.selection === "member") {
  element = <div className="list-group-item">
    <p className="">{props.name}{props.class}</p>
  </div>
  }

  return (
    <div>
      {element}
    </div>
  )
}

export default Collection;