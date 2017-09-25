class App extends React.Component {
  constructor() {
    super();
    this.state = { persons: [] };
  }

  componentDidMount() {
    fetch('https://jsonplaceholder.typicode.com/users')
      .then(data => data.json())
      .then(persons => this.setState({ persons }));
  }

  render() {
    return (
      <div className='pure-u'>
        <PageTitle title='Persons' />
        <Table data={this.state.persons} />
      </div>
    );
  }
}

let PageTitle = (props) =>
<h1>{props.title}</h1>;

let Table = (props) =>
<table className='pure-table'>
  <thead>
    <tr>
      <td>ID</td>
      <td>Name</td>
      <td>Email</td>
      <td>Website</td>
    </tr>
  </thead>
  <tbody>
    {props.data.map((person) =>
      <TableRow key={person.id} data={person} />)}
  </tbody>
</table>;

let TableRow = (props) =>
<tr>
  <td className='id-cell'>{props.data.id}</td>
  <td>{props.data.name}</td>
  <td>{props.data.email}</td>
  <td><a href={props.data.website}>{props.data.website}</a></td>
</tr>;

ReactDOM.render(<App />, document.getElementById('root'));