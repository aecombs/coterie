import React, { useState } from 'react';
import { BrowserRouter as Router, Switch, Route, Link } from 'react-router-dom';
import Announcements from './Announcements';
import Events from './Events';
import Holidays from './Holidays';
import Scriptures from './Scriptures';
import Chapters from './Chapters';

const Dashboard = () => {
  return (
    <Router>
    <div className="">
      <nav className="">
        <ul className="">
          <li><Link to={'/dashboard'} className="">Dashboard</Link></li>
          <li><Link to={'/profile'} className="">Profile</Link></li>
          <li><Link to={'/scriptures'} className="">Scriptures</Link></li>
          <li><Link to={'/holidays'} className="">Holidays</Link></li>
          <li><Link to={'/events'} className="">events</Link></li>
          <li><Link to={'/announcements'} className="">Announcements</Link></li>
        </ul>
      </nav>
      
      <Switch> 
        <Route exact path='/dashboard' component={Dashboard} />
        <Route 
          path={'/'}
          render={() => (
          <Announcements 
          />
        )} />
      </Switch>
    </div>
  </Router>
  )
}

export default Dashboard;