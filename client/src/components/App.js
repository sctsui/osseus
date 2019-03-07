import React from 'react'
import Header from './Header';
import PluginApp from './Plugin/PluginApp';
import GeneratorApp from './Generator/GeneratorApp';
import ProjectSelection from './Project-Selection/ProjectSelection';
import { BrowserRouter, Route} from 'react-router-dom';

class App extends React.Component { 
  render() {
    return (
      <div>
        <BrowserRouter>
          <div>
            <Header/>
            <Route path="/" exact component={ProjectSelection} />
            <Route path="/PluginApp" exact component={PluginApp} />
            <Route path="/GeneratorApp" exact component={GeneratorApp} />
          </div>
        </BrowserRouter>  
      </div>
    );
  }
}
export default App;


