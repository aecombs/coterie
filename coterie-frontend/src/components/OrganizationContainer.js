import React, { useState, useEffect } from 'react';
import Organization from './Organization';
import axios from 'axios';


const OrganizationContainer = (props) => {
  //TODO: logic that checks the URL... is it just the dashboard? then show xyz stuff.

  const [organizationList, setOrganizationList] = useState(null);

  const url = `http://localhost:3000/users/${props.userID}/organizations/`
  useEffect(() => {
    axios.get(url)
      .then( (response) => {
        const currentOrgs = response.data;
        setOrganizationList(currentOrgs);
      })
      .catch((err) => {
        console.log(`There was an error: ${err["message"]}`)
      });
  },[url])


  let orgComponents = undefined
  if (organizationList !== null && organizationList.length > 0) {
    orgComponents = organizationList.map((org) => {
    return(<Organization
      key = {org.id}
      id = {org.id}
      name = {org.name}
      missionStatement = {org.mission_statement}
      totalFunds = {org.total_funds}
      userID = "1"
    />) })} else {
    orgComponents = <p className="open-sans">It looks like you haven't added an org yet!</p>
  }



  return (
    <div className="d-flex container">
      {orgComponents}
    </div>
  )
}

export default OrganizationContainer;