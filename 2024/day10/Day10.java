import java.util.ArrayDeque;
import java.util.ArrayList;
import java.util.Arrays;
import java.util.HashSet;

class Day10 {
    public static void main(String[] args) {
        if (args.length == 0) {
            System.out.println("Hello World!");
        } else {
            System.out.println("Hello " + args[0]);
        }

        var example = "89010123\n" + //
                "78121874\n" + //
                "87430965\n" + //
                "96549874\n" + //
                "45678903\n" + //
                "32019012\n" + //
                "01329801\n" + //
                "10456732";

        var example2 = "89010123\n" + //
                "78121874\n" + //
                "87430965\n" + //
                "96549874\n" + //
                "45678903\n" + //
                "32019012\n" + //
                "01329801\n" + //
                "10456732";

        var inputTxt = "432109865210212123765432101234321098543289654320132112121058\n" + //
                "045678774324301012892343023445456787650198763013241001034569\n" + //
                "187678789465692321001056014896234986456787012894653212123678\n" + //
                "296589921056789433217837895687145675323891233765784589238987\n" + //
                "345437835434576544786921278761010014210710321212098676521067\n" + //
                "032126546323465435695430789760121223121653450303145125430678\n" + //
                "123010567810156543212345699859834321056544067654236012321589\n" + //
                "543213498987657665401030787348765430187432198765987622345432\n" + //
                "654100332394348972342321895201256589196343089543212331056741\n" + //
                "789011241003238981089400776100343678015434567630105449879870\n" + //
                "296721256210169895676510385011892349101325678921256756768987\n" + //
                "129830787323456765410321294332761058210012310123890891057610\n" + //
                "056745698234556786329454301245656567341110567894781232346521\n" + //
                "145894510149645699438765892398305678956923498965654343765430\n" + //
                "236586789838732388454326765567214307967845697874505652894321\n" + //
                "105675676545321267565810674354303212875430786543216701678912\n" + //
                "234321501656130054278989983289432120123421803403545810787600\n" + //
                "321030432567032123123678100176563018987578912012932921298541\n" + //
                "892349803498145031034563210327898101879647810123871092567432\n" + //
                "785056712387236567687654389410787632768756303294562783458943\n" + //
                "176120987656547858998545076585894583459843214785103698327874\n" + //
                "015431234543218947657898145894983298708941025672234567016761\n" + //
                "329122345692105438941763236783470165617652912341013053205430\n" + //
                "478031001785676323430154100102565674320543805432332122124321\n" + //
                "567649872434985610121894321211056589211230796901440345034234\n" + //
                "456659763323810765456765894323987401108921687878981236565105\n" + //
                "306778654310329876365498765012342322317834510967892387156076\n" + //
                "219865011078478901278321001231451015436123423456785493087189\n" + //
                "652104102569560110669125198340760896895034587655476334196892\n" + //
                "765233243458721321701034567654878987764105694344301243234561\n" + //
                "894310321029832459852567878723965430653298743213210358765410\n" + //
                "132123478010741069743478989014436321541056543401821569898324\n" + //
                "098034569123658978654321876101521087632347812114981678876543\n" + //
                "107765678034567867569270965437698794545938903003470549987432\n" + //
                "256872345621098654378187602348967003897821094012561232789501\n" + //
                "345901436438767789210094511059854112766123285723034341076521\n" + //
                "217894387589656231234543223456743245675054176894125652112430\n" + //
                "306105498678543140567672100145101230984169065765898763203431\n" + //
                "495218321067012056478981041234230121243078434965235678976521\n" + //
                "584349452652100987329891230765345698732154567874143454989210\n" + //
                "673458763643211011010010049874556781235463456963056763474321\n" + //
                "567647601781012567892102156743765470346322161012369812565232\n" + //
                "498678432692123476543103095652834387457210052623871001450143\n" + //
                "304509543543001989698234589501921098768921106780982341019898\n" + //
                "213219601982132670787825676501432349810123235691987432870767\n" + //
                "894348732676544561236910787432321056769894344302346549961251\n" + //
                "765210145690125650345210097899867892110765654219854678450340\n" + //
                "890100126780034743094303126934786543023234565678765012321231\n" + //
                "765987034621129802185412235025696541032167876789874349876012\n" + //
                "876856541234988012276543384110567832249054965694101256778123\n" + //
                "965987650945876543543215493201378980158345434783450126789894\n" + //
                "457871056876067875456906780110234589267210321692569034670765\n" + //
                "320432347780128965307878767820199674307890160541078765521254\n" + //
                "011876548991234534218349856936788765216543254332112340432345\n" + //
                "432965432781049621029256743245215656325321067210003451201056\n" + //
                "547876501632898756540178652101304567101452198760116764342767\n" + //
                "656983432542765987438769789012453898212968765641985895433898\n" + //
                "898792323101874104329054210589562456703879454332076016924567\n" + //
                "125601017652963265012123323676571329894312303549165327810430\n" + //
                "034340178943012378901012334568980016765601212678234456901321";

        var input = Day10.parse(inputTxt);
        // System.out.println(input.sumTrailheadScores());
        System.out.println("Trailhead ratings %s".formatted(input.sumTrailheadRatings()));
    }

    public record Input(int[][] topoMap) {
        @Override
        public String toString() {
            StringBuilder sb = new StringBuilder();
            sb.append("Input:\n");
            for (int[] row : topoMap) {
                for (int value : row) {
                    sb.append(value).append(" ");
                }
                sb.append("\n");
            }
            return sb.toString();
        }

        public int sumTrailheadScores() {
            var ths = this.trailheads();
            var sum = 0;
            for (Coord th : ths) {
                var reachableTrailEnds = this.reachableTrailEnds(th);
                sum += reachableTrailEnds.size();
            }
            return sum;
        }

        public int sumTrailheadRatings() {
            var ths = this.trailheads();
            var sum = 0;
            for (Coord th : ths) {
                var pathsToEnd = this.pathsToTrailEnds2(th);
                sum += pathsToEnd.size();
            }
            return sum;
        }

        public HashSet<Coord> trailheads() {
            HashSet<Coord> th = new HashSet<>();
            for (int y = 0; y < this.topoMap.length; y++) {
                for (int x = 0; x < this.topoMap[y].length; x++) {
                    if (topoMap[y][x] == 0) {
                        th.add(new Coord(x, y));
                    }
                }
            }
            System.out.println(th);
            System.out.printf("trailheads: %s%n", th.size());
            return th;
        }

        public HashSet<Coord> reachableTrailEnds(Coord c) {
            var nextSteps = this.reachableNextStep(c);
            var ends = new HashSet<Coord>();
            for (Coord ns : nextSteps) {
                if (this.topoMap[ns.y][ns.x] == 9) {
                    ends.add(ns);
                } else {
                    ends.addAll(this.reachableTrailEnds(ns));
                }
            }
            return ends;
        }

        public ArrayList<TrailPath> pathsToTrailEnds(Coord c) {
            var nextSteps = this.reachableNextStep(c);
            var paths = new ArrayList<TrailPath>();
            for (Coord ns : nextSteps) {
                if (this.topoMap[ns.y][ns.x] == 9) {
                    var path = new ArrayList<Coord>();
                    path.add(ns);
                    path.add(c);
                    paths.add(new TrailPath(path));
                } else {
                    var pathsToEnd = this.pathsToTrailEnds(c);
                    for (TrailPath path : pathsToEnd) {
                        path.coords.add(c);
                    }
                    paths.addAll(pathsToEnd);
                }
            }
            return paths;
        }

        ArrayList<TrailPath> pathsToTrailEnds2(Coord c) {
            ArrayDeque<TrailPath> incompletePaths = new ArrayDeque<>();
            ArrayList<TrailPath> completePaths = new ArrayList<>();

            var initialPath = new TrailPath(new ArrayList<>());
            initialPath.coords().add(c);
            incompletePaths.push(initialPath);

            while (!incompletePaths.isEmpty()) {
                var incompletePath = incompletePaths.pop();
                var lastCoord = incompletePath.coords().getLast();
                var nextSteps = this.reachableNextStep(lastCoord);

                for (var ns : nextSteps) {
                    // Create a new path with copied coordinates
                    var newPath = new TrailPath(new ArrayList<>(incompletePath.coords()));
                    newPath.coords().add(ns);

                    if (this.topoMap[ns.y][ns.x] == 9) {
                        completePaths.add(newPath);
                    } else {
                        incompletePaths.push(newPath);
                    }
                }
            }
            return completePaths;
        }

        public HashSet<Coord> reachableNextStep(Coord c) {
            var reachable = new HashSet<Coord>();
            var current = this.topoMap[c.y][c.x];
            if (c.x - 1 >= 0 && c.x - 1 < this.topoMap[c.y].length && this.topoMap[c.y][c.x - 1] == current + 1) {
                reachable.add(new Coord(c.x - 1, c.y));
            }
            if (c.x + 1 >= 0 && c.x + 1 < this.topoMap[c.y].length && this.topoMap[c.y][c.x + 1] == current + 1) {
                reachable.add(new Coord(c.x + 1, c.y));
            }
            if (c.y - 1 >= 0 && c.y - 1 < this.topoMap.length && this.topoMap[c.y - 1][c.x] == current + 1) {
                reachable.add(new Coord(c.x, c.y - 1));
            }
            if (c.y + 1 >= 0 && c.y + 1 < this.topoMap.length && this.topoMap[c.y + 1][c.x] == current + 1) {
                reachable.add(new Coord(c.x, c.y + 1));
            }
            return reachable;
        }
    }

    record Coord(int x, int y) {
    }

    record TrailPath(ArrayList<Coord> coords) {
    }

    public static Input parse(String input) {
        ArrayList<ArrayList<Integer>> list = new ArrayList<>();
        for (String line : input.split("\n")) {
            if (line.length() == 0) {
                continue;
            }
            boolean hasNonDigit = line.chars()
                    .anyMatch(c -> !Character.isDigit((char) c));

            if (hasNonDigit) {
                char invalidChar = line.chars()
                        .mapToObj(c -> (char) c)
                        .filter(c -> !Character.isDigit(c))
                        .findFirst()
                        .get();
                throw new IllegalArgumentException("String contains non-digit character: " + invalidChar);
            }
            ArrayList<Integer> row = line.chars()
                    .map(c -> Character.getNumericValue(c))
                    .collect(ArrayList::new, ArrayList::add, ArrayList::addAll);
            list.add(row);
        }

        int[][] array = new int[list.size()][list.get(0).size()];

        for (int i = 0; i < list.size(); i++) {
            for (int j = 0; j < list.get(i).size(); j++) {
                array[i][j] = list.get(i).get(j);
            }
        }

        return new Input(array);
    }

}
