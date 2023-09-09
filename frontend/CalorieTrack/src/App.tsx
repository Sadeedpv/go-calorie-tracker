import StripedTable from "./components/StripedTable";
import React from "react";
import Button from "react-bootstrap/Button";
import Modal from "react-bootstrap/Modal";
import { useCalorieStore } from "./utils/useStore";

function App(): JSX.Element {
  const { calorieData, totalCalories, fetchCalorieData, fetchTotalCalories } =
    useCalorieStore();

  const [food, setFood] = React.useState<string | null>('');
  const [info, setInfo] = React.useState<number | null>(null);
  const [calorie, setCalorie] = React.useState<number | null>();

  // React Modal

  const [show, setShow] = React.useState<boolean>(false);

  // React form functions

  const handleShow = () => setShow(true);
  const handleClose = () => setShow(false);

  const handleSubmit = (e: React.FormEvent) => {
    e.preventDefault();
    setShow(false);
    if (food !== '' && typeof(calorie) === "number") {
      // POST the data
      fetch(`${import.meta.env.VITE_PORT}/calories`, {
        method: "POST",
        headers: {
          "Content-Type": "application/json",
        },
        body: JSON.stringify({
          food: food,
          calorie: calorie,
        }),
      })
        .then((res) => res.json())
        .then(() => {
          fetchCalorieData();
          fetchTotalCalories();
        });
    }
  };

  const handleDelete = () => {
    fetch(`${import.meta.env.VITE_PORT}/calories`, {
      method: "DELETE",
      headers: {
        "Content-Type": "application/json",
      },
    })
      .then((res) => res.json())
      .then(() => {
        fetchCalorieData();
        fetchTotalCalories();
      });
  };

  React.useEffect(() => {
    fetchCalorieData();
    fetchTotalCalories();
  }, [fetchCalorieData, fetchTotalCalories]);

  return (
    <div className="container">
      {calorieData && totalCalories ? (
        <StripedTable />
      ) : (
        <p className="loading">Nothing here, Add items 👇👇</p>
      )}

      <div className="buttonContainer">
        <Button variant="primary" onClick={handleShow}>
          Insert here
        </Button>
        <Button variant="primary" onClick={handleDelete}>
          Remove All
        </Button>
      </div>

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
                value={food?.toString()}
                onChange={(e: React.ChangeEvent<HTMLInputElement>) => {
                  setInfo(null)
                  setFood(e.target.value);
                }}
              />
              <Button
                variant="info"
                type="button"
                onClick={() => {
                  try {
                    fetch(
                      `https://api.calorieninjas.com/v1/nutrition?query=${food}`,
                      {
                        method: "GET",
                        headers: {
                          "X-Api-Key": import.meta.env.VITE_API_KEY,
                        },
                      }
                    )
                      .then((res) => res.json())
                      .then((data) => setInfo(data?.items[0]?.calories));
                  } catch (err) {
                    console.log(err);
                  }
                }}
              >
                Get Help
              </Button>
              {info? (
                <p>
                  {food} has {info} calories
                </p>
              ) : (
                <></>
              )}
              <label htmlFor="calorie">Calorie</label>
              <input
                type="number"
                min="0"
                className="form-control"
                id="calorie"
                value={calorie?.toString()}
                placeholder="Enter Calorie"
                onChange={(e: React.ChangeEvent<HTMLInputElement>) => {
                  setCalorie(parseInt(e.target.value));
                }}
              />
            </div>
          </Modal.Body>
          <Modal.Footer>
            <Button variant="primary" type="submit"
              disabled={food === '' || typeof(calorie) !== "number"}
            >
              Save Changes
            </Button>
          </Modal.Footer>
        </form>
      </Modal>
    </div>
  );
}

export default App;
