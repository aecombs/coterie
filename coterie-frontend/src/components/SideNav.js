import React from 'react';
import { BrowserRouter as Router, Switch, Route, Link } from 'react-router-dom';
import OrganizationContainer from './Dashboard/OrganizationContainer';
import Profile from './SideNavElements/Profile';
import Announcements from './SideNavElements/Announcements';
import Events from './SideNavElements/Events';
import Holidays from './SideNavElements/Holidays';
import Scriptures from './SideNavElements/Scriptures';

const SideNav = (props) => {
  const organizationID = props.orgID
  return (
    <Router>
    <div className="d-flex">
      <div className="side-nav">
        <nav className="side-nav">
          <ul className="nav flex-column">
            <li className="nav-item nav-link text-decoration-none"><Link to={'/dashboard'} className="">Dashboard</Link></li>
            <li className="nav-item nav-link"><Link to={'/dashboard/profile'} className="">Profile</Link></li>
            <li className="nav-item nav-link"><Link to={'/dashboard/scriptures'} className="">Scriptures</Link></li>
            <li className="nav-item nav-link"><Link to={'/dashboard/holidays'} className="">Holidays</Link></li>
            <li className="nav-item nav-link"><Link to={'/dashboard/events'} className="">Events</Link></li>
            <li className="nav-item nav-link"><Link to={'/dashboard/announcements'} className="">Announcements</Link></li>
          </ul>
        </nav>
      </div>
      <div className="dashboard-display">
        <Switch>
        {/* <Route exact path='/dashboard' component={Organization} /> */}
        <Route 
            exact path='/dashboard' 
            render={(props) => (
              <OrganizationContainer {...props}
              userID={"1"}
              />
            )}
          />

          <Route path='/dashboard/profile'
           render={(props) => (
            <Profile {...props} 
            userID={"1"}
            />
          )} />
          <Route 
            path='/dashboard/announcements' 
            render={(props) => (
              <Announcements {...props} 
              orgID={"1"}
              />
            )}
          />
          <Route 
            path='/dashboard/events' 
            render={(props) => (
              <Events {...props} 
              orgID={"1"}
              />
            )}
          />
          
          <Route 
            path='/dashboard/holidays' 
            render={(props) => (
              <Holidays {...props} 
              orgID={"1"}
              />
            )}
          />
          <Route 
            path='/dashboard/scriptures' 
            render={(props) => (
              <Scriptures {...props} 
              orgID={"1"}
              />
            )}
          />
        </Switch>
      </div>
    </div>
  </Router>
  )
}

export default SideNav;