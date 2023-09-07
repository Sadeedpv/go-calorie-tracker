import StripedTable from "./components/StripedTable"
import React from 'react'
import Button from 'react-bootstrap/Button'
import Modal from 'react-bootstrap/Modal'
import { calorieData, totalCalories } from "./utils/type";

function App():JSX.Element{
  const [calorieData, setCalorieData] = React.useState<calorieData | null>();
  const [totalCalories, settotalCalories] = React.useState<totalCalories | null>();
  
  const [food, setFood] = React.useState<string | null>()
  const [calorie, setCalorie] = React.useState<number | null>()

  // React Modal

  const [show, setShow] = React.useState<boolean>(false)
  
  // React form functions

  const handleShow = () => setShow(true)
  const handleClose = () => setShow(false)

  const handleSubmit = (e: React.FormEvent) => {
    e.preventDefault()
    if (food && calorie) {
      // POST the data
      fetch(`${import.meta.env.VITE_PORT}/calories`, {
        method: 'POST',
        headers: {
          'Content-Type': 'application/json'
        },
        body: JSON.stringify({
          food: food,
          calorie:calorie
        })
      })
        .then(res => res.json())
      .then(data => console.log(data))
    }
  }

  React.useEffect(() => {
    const fetchCalories = () => {
      fetch(`${import.meta.env.VITE_PORT}/calories`)
        .then((res) => res.json())
        .then((data) => setCalorieData(data));
    };

    const fetchTotalcalories = () => {
      fetch(`${import.meta.env.VITE_PORT}/totalcalories`)
        .then((res) => res.json())
        .then((data) => settotalCalories(data.TotalCalories));
    };

    fetchCalories();
    fetchTotalcalories();
  }, []);

  return (
    <div className="container">
      {calorieData && totalCalories ? (
        <StripedTable calorieData={calorieData} totalCalories={totalCalories} />
      ) : (
        <>Loading...</>
      )}
      <div className="buttonContainer">
        <Button variant="primary" onClick={handleShow}>
          Insert here
        </Button>{" "}
        <Modal show={show} onHide={handleClose} animation={false}>
          <Modal.Header closeButton>
            <Modal.Title>Add today's diet</Modal.Title>
          </Modal.Header>
          <form onSubmit={handleSubmit}>
          <Modal.Body>
              <div className="form-group">
                <label htmlFor="food">Food</label>
                <input
                  type="text"
                  className="form-control" 
                  id="food"
                  placeholder="Enter food"
                  onChange={(e:React.ChangeEvent<HTMLInputElement>) => {
                    setFood(e.target.value)
                  }}
                />
                <label htmlFor="calorie">Calorie</label>
                <input
                  type="number"
                  min='0'
                  className="form-control" 
                  id="calorie"
                  placeholder="Enter Calorie"
                  onChange={(e:React.ChangeEvent<HTMLInputElement>) => {
                    setCalorie(parseInt(e.target.value))
                  }}
                />
              </div>
          </Modal.Body>
          <Modal.Footer>
            <Button variant="primary"  type="submit">
              Save Changes
            </Button>
          </Modal.Footer>
          </form>
        </Modal>
      </div>
    </div>
  );
}

export default App
