import React, {useState, useEffect} from 'react';
import axios from 'axios';
import Newsletter from './Newsletter';
import NewsletterForm from './NewsletterForm';


const Newsletters = (props) => {
  const [newslettersList, setNewslettersList] = useState(null);
  const [visibility, setVisibility] = useState(false);
  const [errorMessage, setErrorMessage] = useState(null);

  const orgID = sessionStorage.getItem('orgID');

  const url = `${process.env.REACT_APP_API_BASE_URL}/users/${props.userID}/organizations/${orgID}/newsletters`

  const updateURL = `${process.env.REACT_APP_API_BASE_URL}/newsletters`

  const addNewsletter = (annObj) => {
    //remove unnecessary id property
    delete annObj["id"];

    axios.post(url, annObj)
    .then((response) => {
      setErrorMessage(`Newsletter ${annObj["header"]} added`);
      window.location.reload();
    })
    
    .catch((error) => {
      setErrorMessage(error.message);
      console.log(`Unable to add newsletter: ${errorMessage}`);
    })
  }

  const updateNewsletter = (annObj) => {
    axios.put(`${updateURL}/${annObj.id}`, annObj)
    .then((response) => {
      setErrorMessage(`Newsletter ${annObj["header"]} was updated`);
      window.location.reload();
    })
    
    .catch((error) => {
      setErrorMessage(error.message);
      console.log(`Unable to update newsletter: ${errorMessage}`);
    })
  }

  const deleteNewsletter = (annID) => {
    axios.delete(`${updateURL}/${annID}`)
    .then((response) => {
      setErrorMessage(`Newsletter ${annID["header"]} was deleted`);
      window.location.reload();
    })
    
    .catch((error) => {
      setErrorMessage(error.message);
      console.log(`Unable to delete newsletter: ${errorMessage}`);
    })
  }

   //toggle visibility of newsletter form component
   const toggleFormVisibility = () => {
    setVisibility(!visibility);
    return;
  }

  //Get all newsletters
  useEffect(() => {
    axios.get(url)
      .then( (response) => {
        const list = response.data;
        setNewslettersList(list);
      })
      .catch((error) => {
        setErrorMessage(error);
        console.log(`There was an error retrieving newsletters: ${error}`)
      });
  },[url])

  let newsletterComponents = undefined
  if (newslettersList !== null && newslettersList.length > 0) {
    newsletterComponents = newslettersList.map((ann) => {
    return(
      <Newsletter
      key={ann.id}
      id={ann.id}
      header={ann.header}
      description={ann.description}
      date={ann.date}
      orgID={ann.organization_id}
      submitNewsletterCallback={updateNewsletter}
      deleteNewsletterCallback={deleteNewsletter}
      />
    )
  })
}


  return (
    <div className="container">
      <div className="d-flex list-group">
        <div className="d-flex py-2 justify-content-between">
          <h4>Newsletters</h4>
          <button className="btn btn-secondary" onClick={toggleFormVisibility}>{ visibility ? "-" : "+"}</button>
        </div>
        <p className={ newsletterComponents !== undefined ? "hidden" : "open-sans" }>You haven't created any newsletters yet.</p>
        <NewsletterForm 
        orgID={orgID}
        visibility={visibility}
        submitNewsletterCallback={addNewsletter}
        onSubmitCallback={toggleFormVisibility}
        />
        {newsletterComponents}
      </div>
    </div>
  )
}

export default Newsletters;