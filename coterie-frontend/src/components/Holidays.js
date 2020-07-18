import React, {useState, useEffect} from 'react';
import axios from 'axios';
import Holiday from './Holiday';

const Holidays = (props) => {
  const [holidaysList, setHolidaysList] = useState(null);

  const url = `http://localhost:3000/users/${props.userID}/organizations/${props.orgID}/holidays`

  useEffect(() => {
    axios.get(url)
      .then( (response) => {
        const list = response.data;
        setHolidaysList(list);
      })
      .catch((error) => {
        console.log(`There was an error retrieving holidays: ${error}`)
      });
  },[url])

  let holidayComponents = undefined
  if (holidaysList !== null && holidaysList.length > 0) {
    holidayComponents = holidaysList.map((hol) => {
    return(
      <Holiday
      key={hol.id}
      id={hol.id}
      name={hol.name}
      date={hol.date}
      description={hol.description}
      orgID={hol.organization_id}
      />
    )
  })
} else {
  holidayComponents = <p className="open-sans">Document your holidays here!</p>
}


  return (
    <div className="container">
      <h4 className="text-left w-100">All Holidays</h4>
      <div className="d-flex list-group">
        {holidayComponents}
      </div>
    </div>
  )
}

export default Holidays;