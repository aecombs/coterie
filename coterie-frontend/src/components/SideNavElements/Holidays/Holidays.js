import React, {useState, useEffect} from 'react';
import axios from 'axios';
import Holiday from './Holiday';
import HolidayForm from './HolidayForm';

const Holidays = (props) => {
  const [holidaysList, setHolidaysList] = useState(null);
  const [visibility, setVisibility] = useState(false);
  const [errorMessage, setErrorMessage] = useState(null);

  const orgID = sessionStorage.getItem('orgID');

  const url = `${process.env.REACT_APP_API_BASE_URL}/users/${props.userID}/organizations/${orgID}/holidays`

  const updateURL = `${process.env.REACT_APP_API_BASE_URL}/holidays`


  const addHoliday = (holObj) => {
    //remove unnecessary id property
    delete holObj["id"];

    axios.post(url, holObj)
    .then((response) => {
      setErrorMessage(`Holiday ${holObj["name"]} added`);
      window.location.reload();
    })
    
    .catch((error) => {
      setErrorMessage(error.message);
      console.log(`Unable to add holiday: ${errorMessage}`);
    })
  }

  const updateHoliday = (holObj) => {
    axios.put(`${updateURL}/${holObj.id}`, holObj)
    .then((response) => {
      setErrorMessage(`Holiday ${holObj["name"]} was updated`);
      window.location.reload();
    })
    
    .catch((error) => {
      setErrorMessage(error.message);
      console.log(`Unable to add holiday: ${errorMessage}`);
    })
  }

  const deleteHoliday = (holID) => {
    axios.delete(`${updateURL}/${holID}`)
    .then((response) => {
      setErrorMessage(`Holiday ${holID["name"]} was deleted`);
      window.location.reload();
    })
    
    .catch((error) => {
      setErrorMessage(error.message);
      console.log(`Unable to delete holiday: ${errorMessage}`);
    })
  }

   //toggle visibility of holiday form component
   const toggleFormVisibility = () => {
    setVisibility(!visibility);
    return;
  }


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
      submitHolidayCallback={updateHoliday}
      deleteHolidayCallback={deleteHoliday}
      />)
    })
  }


  return (
    <div className="container">
      <div className="d-flex list-group">
        <div className="d-flex py-2 justify-content-between">
          <h4>Holidays</h4>
          <button className="btn btn-secondary" onClick={toggleFormVisibility}>{ visibility ? "-" : "+"}</button>
        </div>
        <p className={ holidayComponents !== undefined ? "hidden" : "open-sans" }>You haven't created any holidays yet.</p>
        <HolidayForm 
        orgID={orgID}
        visibility={visibility}
        submitHolidayCallback={addHoliday}
        onSubmitCallback={toggleFormVisibility}
        />
        {holidayComponents}
      </div>
    </div>
  )
}

export default Holidays;