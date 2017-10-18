import React, {Component} from 'react';
import pkgFetcher from './utils/pkgFetcher'
import Pkg from './Pkg';
import {toast} from 'react-toastify';

export default class PkgBrowser extends Component {
  constructor(props) {
    super(props);

    this.state = {
      pkgRef: 'github.com/opspec-pkgs/git.clean#1.0.0',
    };

    this.openPkg(this.state.pkgRef);
  }

  openPkg(pkgRef) {
    pkgFetcher.fetch(pkgRef)
      .then(pkg => this.setState({pkgRef, pkg}))
      .catch(error => {
        toast.error(error.message);
      });
  }

  handleSubmit(e) {
    e.preventDefault();

    this.openPkg(this.state.pkgRef);
  }

  render() {
    return (
      <div className='container'>
        <div>
          <form onSubmit={e => this.handleSubmit(e)} style={{paddingTop: '25px'}}>
            <div className='form-group'>
          <span className='input-group input-group-lg'>
            <input className='form-control'
                   id='pkgRef'
                   type='text'
                   value={this.state.pkgRef}
                   onChange={e => this.setState({pkgRef: e.target.value})}
                   placeholder="/absolute/path or host/path/git-repo#tag"/>
            <span className='input-group-btn'>
              <button className='btn btn-primary' id='pkgRef_Submit' type='submit'>open</button>
            </span>
          </span>
            </div>
          </form>
          {this.state.pkg ? <Pkg value={this.state.pkg} pkgRef={this.state.pkgRef}/> : null}
        </div>
      </div>
    )
  }

}
