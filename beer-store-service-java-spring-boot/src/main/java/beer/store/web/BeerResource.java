package beer.store.web;

import java.util.List;

import org.springframework.beans.factory.annotation.Autowired;
import org.springframework.web.bind.annotation.DeleteMapping;
import org.springframework.web.bind.annotation.GetMapping;
import org.springframework.web.bind.annotation.PathVariable;
import org.springframework.web.bind.annotation.PostMapping;
import org.springframework.web.bind.annotation.PutMapping;
import org.springframework.web.bind.annotation.RequestBody;
import org.springframework.web.bind.annotation.RequestMapping;
import org.springframework.web.bind.annotation.RequestParam;
import org.springframework.web.bind.annotation.RestController;

import beer.store.model.Beer;
import beer.store.service.BeerService;

@RestController
@RequestMapping("beer")
public class BeerResource {

  @Autowired
  private BeerService service;

  @GetMapping("list")
  public List<Beer> list(@RequestParam(defaultValue = "") String search, @RequestParam(defaultValue = "1") Integer page,
      @RequestParam(defaultValue = "10") Integer pageSize) {
    return service.list(search, page, pageSize);
  }

  @GetMapping("{id}")
  public Beer find(@PathVariable(name = "id", required = true) Integer idBeer) {
    return service.find(idBeer);
  }

  @PostMapping("save")
  public Beer insert(@RequestBody Beer beer) {
    return service.insert(beer);
  }

  @PutMapping("save")
  public Beer update(@RequestBody Beer beer) {
    return service.update(beer);
  }

  @DeleteMapping("{id}")
  public String del(@PathVariable(name = "id", required = true) Integer idBeer) {
    service.del(idBeer);
    return "OK";
  }
}