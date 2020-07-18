import React, { useState, useEffect } from 'react';
import Chapter from './Chapter';
import axios from 'axios';

const Scripture = (props) => {
  const [chaptersList, setChaptersList] = useState(null);

  const url = `http://localhost:3000/organizations/${props.orgID}/scriptures/${props.id}/chapters`

  useEffect(() => {
    axios.get(url)
      .then( (response) => {
        const list = response.data;
        setChaptersList(list);
      })
      .catch((error) => {
        console.log(`There was an retrieving chapters: ${error}`)
      });
  },[url])
  
  let chapterComponents = undefined
  if (chaptersList !== null && chaptersList.length > 0) {
    chapterComponents = chaptersList.map((chap) => {
    return (<Chapter 
          key={chap.id}
          id={chap.id}
          name={chap.name}
          text={chap.text}
          position={chap.position}
          scripID={chap.scripture_id}
        />)
    });


  return (
    <div>
      <p className="open-sans">{props.name}</p>
      {chapterComponents}
      
    </div>
  )
}
}

export default Scripture;