import React, { useState, useEffect } from 'react';
import axios from 'axios';
import Announcement from './Announcement';
import Event from './Event';
import Holiday from './Holidays/Holiday';
import Scripture from './Scripture';
import Member from './Member';

const Collection = (props) => {
  const [elementList, setElementList] = useState([]);
  
  const url = `http://localhost:3000/users/${1}/organizations/${props.orgID}/${props.selection}`

  useEffect(()=>{
    axios.get(url)
      .then((response) => {
        const elements = response.data;
        setElementList(elements);
      })
      .catch((error) => {
        console.log(`There was an error retrieving elements: ${error}`)
      });
  }, [url, props.selection])

  
  const listOfElements = elementList.map((element) => {
    switch (props.selection) {
    case "announcements":
      return(<div className="list-group-item">
        <p className="">{props.header}{props.date}</p>
      </div>);
    
    case "events" || "holidays":
      return(<div className="list-group-item">
        <p className="">{props.name}{props.date}</p>
      </div>);
    
    case "scriptures": 
      return(<div className="list-group-item">
        <p className="">{props.name}</p>
      </div>);
    case "chapters": 
      return(<div className="list-group-item">
        <p className="">{props.name}</p>
      </div>);
    case "members":
      return(<div className="list-group-item">
        <p className="">{props.name}{props.class}</p>
      </div>);

    default:
      return(<div className="list-group-item">
        <p className="">Unable to retrieve {props.selection}</p>
      </div>);
    }
  });


  return (
    <div>
      {listOfElements}
    </div>
  )
}

export default Collection;